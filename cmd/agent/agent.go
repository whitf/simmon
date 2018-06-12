package main

/*
 * simmon-agent.go
 * simmon - simple monitoring
 */

import (
 	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
)

type QoS struct {
	Agent	 					string `json:"agent"`
	AgentAddress 				string `json:"agentAddress"`
	//AgentPort 					string `json:"agentPort"`
	AgentId						string `json:"agentId"`
	AgentVersion				string `json:"agentVersion"`
	Alert 						bool `json:"alert"`
	AlertLevel 					string `json:"alertLevel"`
	AlertValue 					string `json:"alertValue"`
	HeartbeatThreshold 			float64 `json:"heartbeatThreshold"`
	Date 						string `json:"date"`
	Msg 						string `json:"msg"`
	QoSType						string `json:"qosType"`
}

type Service struct  {
	Delay			int `json:"delay"`
	Name			string `json:"name"`
	Options			[]ServiceOption	`json:"options"`
}

type ServiceOption struct {
	Key				string `json:"key"`
	Value 			string `json:"value"`
}

var c Config

func handleQoSFunction(d int, o []ServiceOption, f func(int, []ServiceOption) int) int {
	return f(d, o)
}

//@TODO -> Add the ability (unix socket?  tcp from node?  Both?) for a signal to be passed to
// the agent that will set done to true and shut it down gracefully.

//@TODO -> move most of this into the init function
func main() {
	log.Println("starting simmon-agent - simple monitoring agent")
	log.Println()
	log.Println("reading conf")

	flag.Parse()
	if(!Filename.set) {
		log.Println("Configuration file not set (no --conf <file> option).")
		log.Println("Trying default: /etc/simmon/simmon-agent.conf")
		Filename.value = "/etc/simmon/simmon-agent.conf"
	}

	doReadConfig(Filename.value)

	r, _ := regexp.Compile("(<|>)")
	if r.FindStringIndex(c.Name) != nil {
		log.Println()
		log.Println("The characters '<' and/or '>' are not valid in an agent name.")
		log.Println(" Name: " + c.Name)
		log.Println()
		os.Exit(1)
	}

	log.Println("starting agent monitoring with configuration:")
	log.Println(" app: ", c.App)
	log.Println(" name: ", c.Name)
	log.Println(" version: ", c.Version)
	log.Println(" id: ", c.Uuid)
	log.Println( "heartbeat threshold: " + strconv.FormatFloat(c.HeartbeatThreshold, 'f', -1, 64))
	log.Println(" node(s):")
	log.Println(c.Nodes)
	log.Println(" services:")
	log.Println(c.Services)

	// Load map of nodes to check in with.
	NodeMap = make(map[string]*Node)
	for n := range c.Nodes {
		AddNode(&c.Nodes[n])
	}

	// Load map of services to run for this agent.
	smap := make(map[int]Service)
	for s := range c.Services {
		smap[s] = c.Services[s]
	}

	log.Println()
	log.Println("service map")
	log.Println(smap)

	// Load map of monitoring functions that will be run.
	fmap := map[string]func(int, []ServiceOption) int {
		"memfree": memfree,
		"swapfree": swapfree,
		"process": process,
		"load": load,
		"diskFree": diskFree,
	}

	// load initial thresholds from conf file
	mt = MemoryThreshold {
		memoryFreePercent: 0.25,
		swapFreePercent: 0.25,
	}

	log.Println("startup complete, starting monitoring services")

	// Blocking channel stays open.
	done := make(chan bool, 1)

	log.Println("agent is going online")
	online()
	
	// Kick off monitoring functions on their own thread(s).
	// Except heartbeat, which is special.  (Handled by online function.)
	for service := range smap {
		log.Println(" launched: " + smap[service].Name)
		go handleQoSFunction(smap[service].Delay, smap[service].Options, fmap[smap[service].Name])
	}

	log.Println("all monitoring services launched.")

	<-done

	log.Println("simmon-agent on " + c.Name + " closed")
}

func init() {
	log.SetFlags(log.Lshortfile)
}
