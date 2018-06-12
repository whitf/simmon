package main

/**
 * cmd/agent/netComm.go
 * simmon - simple monitoring
 *
 * Handle network connections to a node or nodes.
 */

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type Node struct {
	Alive 				bool			`json:"alive"`
	Host				string			`json:"host"`
	Port 				string			`json:"port"`
	Pulse 				int				`json:"pulse"`
}

var NodeMap map[string]*Node

func AddNode(n *Node) {
	NodeMap[n.Host] = n
	log.Println("Adding node " + n.Host + ":" + n.Port)
}

func DelNode(n *Node) {
	delete(NodeMap, n.Host)
}

func online() int {
	q := QoS {
		Agent: c.Name,
		AgentAddress: "agent address",
		AgentId: c.Uuid,
		AgentVersion: c.Version,
		Alert: false,
		AlertLevel: "Clean",
		AlertValue: "ONLINE",
		Date: time.Now().String(),
		Msg: "ONLINE",
		QoSType: "ONLINE",
	}

	//var nodeHost, nodePort = "", ""
	for host, n := range NodeMap { 
		controlPort, parseErr := strconv.Atoi(n.Port)
		check(parseErr)
		controlPort = 1 + controlPort

		conn, err := net.Dial(c.Protocol, host + ":" + strconv.Itoa(controlPort))
		connCheck(err)

		enc := gob.NewEncoder(conn)
		enc.Encode(q)
		conn.Close()

		n.Alive = true
		n.resetPulse()
		go n.heartbeat()
	}

	log.Println(" agent " + c.Name + " is online.")
	return 0
}

func (q *QoS) report() int {
	if len(NodeMap) > 0 {
		// send to node(s)
		for host, n := range NodeMap {
			if n.Alive {
				log.Println("reporting " + q.QoSType + " to " + host + ":" + n.Port)

				conn, err := net.Dial(c.Protocol, host + ":" + n.Port)
				if err == nil {
					enc := gob.NewEncoder(conn)
					enc.Encode(q)
					conn.Close()
				}
			}
		}

		return 0
	} else {
		log.Println()
		log.Println("No active nodes.")
		log.Println(" Local qos storage not yet implemented.  Discarding.")
		log.Println()
		return 1
	}
}

func (q *QoS) reportHeartbeat() int {
	//@TODO individual hosts need seperate "pulses" for heartbeats....

	// Report heartbeat to each node.
	for host, n := range NodeMap {
		log.Println("reporting to " + host + ":" + n.Port)
		conn, err := net.Dial(c.Protocol, host + ":" + n.Port)
		
		if err != nil {
			if nerr, ok := err.(net.Error); ok {
				fmt.Println(nerr)
			//if err.connect == "connection refused" {
				log.Println("net error of some sort...")
				log.Println("incrementing heartbeat pulse")
				n.extendPulse()
				n.setPulse()
			} else {
				log.Fatal(err)
			} 
		} else {
			if n.Pulse > 16 {
				n.resetPulse()
			}
			
			enc := gob.NewEncoder(conn)
			enc.Encode(q)
			conn.Close()
		}
		
	}

	return 0
}