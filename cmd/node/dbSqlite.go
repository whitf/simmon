package main

/**
 * cmd/node/sqLite.go
 * simmon - simple monitoring
 *
 * sqlite database functions
 */

import (
 	"database/sql"
 	"log"
 	"os"

 	_ "github.com/mattn/go-sqlite3"
)

func dbSqliteInit() {
	var err error
	db, err = sql.Open("sqlite3", c.DbName)
	if err != nil {
		log.Println("Error opening sqlite3 database connection.")
		log.Println()
		log.Println(err)
		log.Println()

		os.Exit(1)
	}
}

func writeQoSSqlite(q *QoS) int {
	var qosInsert = "INSERT INTO qos (alert, alertLevel, alertValue, agent, agentAddress, agentId, agentVersion, heartbeatThreshold, date, msg, qosType) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	statement, _ := db.Prepare(qosInsert)
	statement.Exec(q.Alert, q.AlertLevel, q.AlertValue, q.Agent, q.AgentAddress, q.AgentId, q.AgentVersion, q.HeartbeatThreshold, q.Date, q.Msg, q.QoSType)

	return 0
}

func writeAgentSqlite(q *Agent) int {
	return 0
}

func removeAgentSqlite(a *Agent) int {
	return 0
}

func writeAlertSqlite(q *QoS) int {
	var stmt = "SELECT COUNT(*) AS count FROM alert WHERE agent=\"" + q.Agent + "\" AND qosType=\"" + q.QoSType + "\" AND deleted=0"
	statement, _ := db.Prepare(stmt)

	var count = 0
	_ = statement.QueryRow().Scan(&count)

	if count < 1 {
		stmt = "INSERT INTO alert (alert, alertLevel, alertValue, agent, agentAddress, agentId, agentVersion, heartbeatThreshold, date, msg, qosType, deleted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"	
		statement, _ = db.Prepare(stmt)
		statement.Exec(q.Alert, q.AlertLevel, q.AlertValue, q.Agent, q.AgentAddress, q.AgentId, q.AgentVersion, q.HeartbeatThreshold, q.Date, q.Msg, q.QoSType, false)
	}
	
	return 0
}

func clearDbAlertSqlite(aName string, qType string) int {
	var stmt = "UPDATE alert SET deleted=true WHERE agent=\"" + aName + "\" AND qosType=\"" + qType + "\""

	statement, _ := db.Prepare(stmt)
	statement.Exec()

	return 0
}

func loadAlertsSqlite() map[string]*QoS {
	alertsFromDatabase := make(map[string]*QoS)

	//var stmt = "SELECT "




	return alertsFromDatabase
}

func loadAgentListSqlite() map[string]*Agent {
	agentsFromDatabase := make(map[string]*Agent)
	return agentsFromDatabase
}

