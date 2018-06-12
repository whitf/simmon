package main

/**
 * agent/heartbeat.go
 * heartbeat monitoring function
 */

import (
	"time"
)

var pulse int

func (n *Node) setHeartbeatPulse() {
	pulse = n.Pulse
}

func (n *Node) setPulse() {
	n.Pulse = pulse
}

func (n *Node) resetPulse() {
	n.Pulse = 15
	pulse = 15
}

// in case of failed heartbeat from agent -> node
// scale up the wait time between heartbeats up to 15 minutes
func (n *Node) extendPulse() {
	if pulse > 900 {
		pulse = 901
	} else {
		pulse = pulse * 2
	}
}

func (n *Node) heartbeat() {
	for {
		q := QoS {
			Alert: false,
			AlertLevel: "Clean",
			AlertValue: "thum-thump",
			Agent: c.Name,
			AgentAddress: "agent-address-goes-here",
			AgentId: c.Uuid,
			AgentVersion: c.Version,
			HeartbeatThreshold: 30.0,
			Date: time.Now().String(),
			Msg: "HEARTBEAT",
			QoSType: "HEARTBEAT",
		}

		q.reportHeartbeat()
		
		time.Sleep(time.Duration(pulse) * time.Second)
	}
}

