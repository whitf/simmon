package main

/**
 * node/alerts.go
 * simmon - simple monitoring
 */

import (
 	"log"
 	"sync"
 	"time"
)

var (
	AlertLock sync.Mutex
	Alerts map[string]*QoS
)

func addAlert(aName string, qType string, q *QoS) int {

	key := aName + "<THROWING-->" + qType

	AlertLock.Lock()
	Alerts[key] = q
	AlertLock.Unlock()

	log.Println("calling writeAlert(q) for")
	log.Println(q)

	return writeAlert(q)
}

func delAlert(aName string, qType string) int {

	key := aName + "<THROWING-->" + qType

	delete(Alerts, key)

	// soft delete from database
	return clearDbAlert(aName, qType)
}


// Generic functions that pass off to more specific alert handlers.
func doHandleAlert(q *QoS) int {

	return addAlert(q.Agent, q.QoSType, q)
}

func doClearAlert(aName string, qType string) int {

	if _, ok := Alerts[aName + "<THROWING-->" + qType]; ok {
		return delAlert(aName, qType)
	}

	return 0
}

func handleHeartbeatAlert(a *Agent) int {
	a.markOffline()

	q := QoS {
		Alert: true,
		AlertLevel: "CRITICAL",
		AlertValue: "Heartbeat Alert",
		Agent: a.Name,
		AgentAddress: a.Address,
		AgentId: a.Name,
		AgentVersion: a.Version,
		HeartbeatThreshold: a.HeartbeatThreshold,
		Date: time.Now().String(),
		Msg: "HEARTBEAT",
		QoSType: "HEARTBEAT",
	}

	return addAlert(q.Agent, q.QoSType, &q)
}

func (a *Agent) clearHeartbeatAlert() int {
	a.markOnline()
	return doClearAlert(a.Name, "HEARTBEAT")
}

var writeAlert = func(q *QoS) int {
	return 0
}

var clearDbAlert = func(a string, q string) int {
	return 0
}

var loadAlerts = func() map[string]*QoS {
	return make(map[string]*QoS)

}
