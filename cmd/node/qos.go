package main

/**
 * cmd/node/qos.go
 * simmon - simple monitoring
 *
 * Structs and functions to handle QoS data.
 */

type QoS struct {
	Alert 						bool `json:"alert"`
	AlertLevel 					string `json:"alertLevel"`
	AlertValue 					string `json:"alertValue"`
	Agent	 					string `json:"agent"`
	AgentAddress 				string `json:"agentAddress"`
	//AgentPort 					string `json:"agentPort"`
	AgentId						string `json:"agentId"`
	AgentVersion				string `json:"agentVersion"`
	HeartbeatThreshold 			float64 `json:"heartbeatThreshold"`
	Date 						string `json:"date"`
	Msg 						string `json:"msg"`
	QoSType						string `json:"qosType"`
}

var writeQoS = func(q *QoS) int {
	// Global placeholder.
	return 0
}

func handleQoS(q *QoS) int {

	go writeQoS(q)

	if q.QoSType == "HEARTBEAT" {
		return handleHeartbeat(q)
	} else {
		// Handle self-reporting alerts.
		if q.Alert {
			return doHandleAlert(q)
		} else {
			return doClearAlert(q.Agent, q.QoSType)
		}
	}

}