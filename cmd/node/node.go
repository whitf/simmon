package main

/**
 * simmon-node.go
 * simmon - simple monitoring
 */

import (
 	"database/sql"
 	"encoding/gob"
 	"flag"
	"log"
	"net"
	"strconv"
	"time"
)

var c Config
var db *sql.DB

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	q := &QoS{}
	dec.Decode(q)

	q.AgentAddress = conn.RemoteAddr().String()
	
	handleQoS(q)
	conn.Close()
}

func monControl() int {

	monPort, parseErr := strconv.ParseInt(c.Port, 10, 32)
	if parseErr != nil {
		panic(parseErr)
	}
	monPort = 1 + monPort

	log.Println(" starting monitoring controller on port :" + strconv.FormatInt(monPort, 10))

	ln, err := net.Listen("tcp", ":" + strconv.FormatInt(monPort, 10))
	check(err)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error in Listen.Accept() -> ", err)
			continue
		}

		go handleMonControl(conn)
	}

	return 0
}

func handleMonControl(conn net.Conn) int {
	log.Println("handling monitor control message")

	dec := gob.NewDecoder(conn)
	q := &QoS{}
	dec.Decode(q)
	conn.Close()

	switch q.QoSType {
	case "ONLINE":
		log.Println("marking agent " + q.Agent + " online")

		a := Agent {
			Address: conn.RemoteAddr().String(),
			Alive: true,
			HeartbeatThreshold: 30.0,
			Last: time.Now(),
			Name: q.Agent,
			Version: q.AgentVersion,
		}

		return a.addAgent()
	case "OFFLINE":
		log.Println("marking agent " + q.Agent + " offline")
	}

	return 0
}

func nodeControl(simchan chan string) int {

	for {
		log.Println("top of nodeControl loop")
		log.Println("node control message " + <-simchan)

		time.Sleep(time.Duration(15) * time.Second)
	}

	return 0
}

func main() {
	log.Println("starting app...")
	log.Println("reading conf...")

	flag.Parse()
	if(!Filename.set) {
		log.Println("configuration file not set (no --conf <file> option, trying default: /etc/simmon/simmon-node.conf)")
		Filename.value = "/etc/simmon/simmon-node.conf"
	}

	doReadConfig(Filename.value)

	switch c.Storage {
	case "postgres":
		log.Println("connecting to postgres data store")

		writeQoS = writeQoSPostgres
		writeAlert = writeAlertPostgres
		writeAgent = writeAgentPostgres
		clearDbAlert = clearDbAlertPostgres
		loadAlerts = loadAlertsPostgres

	case "mongod":
		log.Println("connecting to mongod data store")

		writeQoS = func(q *QoS) int {
			log.Println("writing to mongod data store")

			return 0
		}

	case "sqlite":
		log.Println("connecting to sqlite data store")

		writeQoS = writeQoSSqlite
		writeAlert = writeAlertSqlite
		writeAgent = writeAgentSqlite
		clearDbAlert = clearDbAlertSqlite
		loadAlerts = loadAlertsSqlite

		dbSqliteInit()

	default:
		// No database configured.  Print monitoring data.

		log.Println("printing data (no database configured in simmon-node.conf)")
		log.Println()
		log.Println("WARNING - Monitoring data will not be persistant on node restart.")
		log.Println()

		writeQoS = func(q *QoS) int {
			log.Println(q)
			return 0
		}

		writeAlert = func(q *QoS) int {
			log.Println(q)
			return 0
		}

		writeAgent = func (a *Agent) int {
			log.Println(a)
			return 0
		}

		// No entry for clearDbAlert, loadAlertList necessary.
	}

	Alerts = loadAlerts()
	loadAgentsList()

	log.Println("Starting monitoring control module")
	go monControl()

	log.Println("Launching heartbeat loop")
	go heartbeat()

	log.Println("Starting node control channel")
	simchan := make(chan string, 100)
	go nodeControl(simchan)

	log.Println("Starting monitoring listener...")

	ln, err := net.Listen("tcp", ":" + c.Port)
	if err != nil {
		panic(err)
	}

	log.Println(" node listening on: " + c.Listen + ":" + c.Port)
	
	log.Println("Starting node-api listener...")

	go nodeApi()


	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error in Listen.Accept() -> ", err)
			continue
		}

		go handleConnection(conn)
	}
}


func init() {
	log.SetFlags(log.Lshortfile)
}