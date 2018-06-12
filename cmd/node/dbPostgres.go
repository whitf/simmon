package main

/**
 * cmd/node/dbPostgres.go
 * simmon - simple monitoring
 *
 * postgres database functions
 */

 import (
 	//"database/sql"

 	//_ "github.com/lib/pq"
 )


func writeQoSPostgres(q *QoS) int {
	return 0
}

func writeAgentPostgres(a *Agent) int {
	return 0
}

func removeAgentPostgres(a *Agent) int {
	return 0
}

func writeAlertPostgres(q *QoS) int {
	return 0
}

func clearDbAlertPostgres(a string, q string) int {
	return 0
}

func loadAlertsPostgres() map[string]*QoS {
	alertsFromDatabase := make(map[string]*QoS)
	return alertsFromDatabase
}

func loadAgentListPostgres() map[string]*Agent {
	agentsFromDatabase := make(map[string]*Agent)
	return agentsFromDatabase
}

