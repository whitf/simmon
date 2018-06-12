package main

/**
 * node-heartbeat.go
 * simmon - simple monitoring
 *
 * Responsible for keeping track of agents and how recently they have 
 * checked in through a heartbeat message.
 */

import (
 	"log"
 	"os"
 	"time"
)

var alive bool

func heartbeat() int {
	if alive {
		log.Println("only one heartbeat per node")
		os.Exit(1)
	}

	alive = true
	
	// loop through registered agents and compare their last heartbeat check in to their threshold.
	for {

		for a := range Agents {
			// Skip dead agents.
			if Agents[a].Alive {

				diff := time.Now().Sub(Agents[a].Last)
				if diff.Seconds() >= Agents[a].HeartbeatThreshold {
					// Throw heartbeat alert.
					go handleHeartbeatAlert(Agents[a])
				}
			}
		}

		time.Sleep(time.Duration(10) * time.Second)

		log.Println()
		log.Println("alerts")
		log.Println(Alerts)
		log.Println()
		log.Println("agents")
		log.Println(Agents)
		log.Println()
	}

	return 0
}

func handleHeartbeat(q *QoS) int {
	if _, ok := Agents[q.Agent]; ok {
		Agents[q.Agent].thump()
	} else {
		a := q.agentFromHeartbeat()
		return a.addAgent()
	}

	return 0
}

