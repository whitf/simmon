package main

/**
 * cmd/node/nodeApi.go
 * simmon - simple monitoring
 *
 */


import (
 	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func GetAlerts(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - /alerts " + time.Now().String())

	var a []*QoS

	for _, v := range Alerts {
		a = append(a, v)
	}

	json.NewEncoder(w).Encode(a)
}

func GetAgents(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - /agents " + time.Now().String())

	var a []*Agent

	for _, v := range Agents {
		a = append(a, v)
	}

	json.NewEncoder(w).Encode(a)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - /version " + time.Now().String())

	version := map[string]string{"version": c.Version}

	//version := `{"version": "` + c.Version + `"}`
	json.NewEncoder(w).Encode(version)
}


func nodeApi() {

	router := mux.NewRouter()

	router.HandleFunc("/alerts", GetAlerts).Methods("GET")
	router.HandleFunc("/agents", GetAgents).Methods("GET")
	router.HandleFunc("/version", GetVersion).Methods("GET")

	log.Println(" node-api listening on: " + c.ApiListen + ":" + c.ApiPort)

	log.Fatal(http.ListenAndServe(":" + c.ApiPort, router))
}