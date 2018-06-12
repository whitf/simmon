package main

/**
 * node/agent.go
 * simmon - simple monitoring
 *
 */

import (
 	"sync"
	"time"
)

type Agent struct {
	Alive							bool `json:"alive"`
	Address							string `json:"address"`
	HeartbeatThreshold				float64 `json:"heartbeatThreshold"`
	Id 								string `json:"id"`
	Last							time.Time `json:"last"`
	Name							string `json:"name"`
	Port							string `json:"port"`
	Version							string `json:"version"`
}

var (
	AgentsLock sync.Mutex
	Agents map[string]*Agent
)

func (a *Agent) addAgent() int {

	if _ , ok := Agents[a.Name]; ok {
		// Agent is already in the list.  Update with passed in values.
		AgentsLock.Lock()
		Agents[a.Name] = a
		AgentsLock.Unlock()
	} else {
		// Add new agent to monitoring list.
		AgentsLock.Lock()
		Agents[a.Name] = a
		AgentsLock.Unlock()
	}

	a.markOnline()
	return writeAgent(a)
}

func (a *Agent) delete() int {
	AgentsLock.Lock()
	delete(Agents, a.Name)
	AgentsLock.Unlock()

	return deleteAgent(a)
}

func (a *Agent) markOnline() int {
	a.Alive = true
	return 0
}

func (a *Agent) markOffline() int {
	a.Alive = false
	return 0
}

func (a *Agent) thump() int {
	a.Alive = true
	a.Last = time.Now()
	return doClearAlert(a.Name, "HEARTBEAT")
}

func (q *QoS) agentFromHeartbeat() *Agent {
	a := &Agent {
		Alive: true,
		Address: q.AgentAddress,
		HeartbeatThreshold: q.HeartbeatThreshold,
		Id: q.AgentId,
		Last: time.Now(),
		Name: q.Agent,
		Version: q.AgentVersion,
	}

	return a
}

func loadAgentsList() {
	if len(Agents) < 1 {
		Agents = map[string]*Agent{}
	}

}

var writeAgent = func(a *Agent) int {
	// Global placeholder.
	return 0
}

var deleteAgent = func(a *Agent) int {
	// Global placeholder.
	return 0
}