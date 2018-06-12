package main

/**
 * cmd/agent/config.go
 * simmon - simple monitoring
 * 
 * Agent Config structure and functions.
 */


import (
 	"encoding/json"
 	"flag"
 	"io/ioutil"
 	"log"
)

type Config struct {
	App							string					`json:"app"`
	HeartbeatThreshold 			float64					`json:"heartbeatThreshold"`
	Name						string					`json:"name"`
	Nodes						[]Node					`json:"nodes"`
	Port						string 					`json:"port"`
	Protocol					string					`json:"protocol"`
	Services 					[]Service		 		`json:"services"`
	Uuid						string					`json:"uuid"`
	Version 					string 					`json:"version"`
}

type confFlag struct {
	set bool
	value string
}

var Filename confFlag

func init() {
	flag.Var(&Filename, "conf", "simmon agent configuration")
}

//@TODO: add in better error checking (file not found, etc)
func doReadConfig(cfile string) {
	cfg, _ := ioutil.ReadFile(cfile)

	err := json.Unmarshal(cfg, &c)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connCheck(e error) {
	if e != nil {
		log.Fatal("Connection Error agent -> node", e)
	}
}

func (cf *confFlag) Set(x string) error {
	cf.value = x
	cf.set = true
	return nil
}

func (cf *confFlag) String() string {
	return cf.value
}
