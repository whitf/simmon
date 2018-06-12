package main

/**
 * cmd/web/server.go
 * simmon - simple monitoring
 *
 */

import (
 	"encoding/json"
 	"flag"
 	"html/template"
 	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var c Config

type Agent struct {
	Alive							bool `json:"alive"`
	Address							string `json:"address"`
	HeartbeatThreshold				int `json:"heartbeatThreshold"`
	Id 								string `json:"id"`
	Last							time.Time `json:"last"`
	Name							string `json:"name"`
	Port							string `json:"port"`
	Version							string `json:"version"`
}

type Alert struct {
	Alert 			bool `json:"alert"`
	AlertLevel 		string `json:"alertLevel"`
	AlertValue 		string `json:"alertValue"`
	Agent	 		string `json:"agent"`
	AgentId			string `json:"agentId"`
	Date 			string `json:"date"`
	Msg 			string `json:"msg"`
	QoSType			string `json:"qosType"`
}

type PageData struct {
	Version				string
	Alerts	 			[]Alert
	Agents				[]Agent
}

type Version struct {
	Version 			string `json:"version"`

}

func getAgentsFromApi() []Agent {
	var a []Agent

	apiClient := http.Client {
		Timeout: time.Second * 2,
	}

	url := "http://" + c.ApiHost + ":" + c.ApiPort + "/agents"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	check(err)

	res, getErr := apiClient.Do(req)
	check(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	check(readErr)

	jsonErr := json.Unmarshal(body, &a)
	check(jsonErr)

	return a
}

func getAlertsFromApi() []Alert {
	var a []Alert

	apiClient := http.Client {
		Timeout: time.Second * 2,
	}

	url := "http://" + c.ApiHost + ":" + c.ApiPort + "/alerts"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	check(err)

	res, getErr := apiClient.Do(req)
	check(getErr)

	body, readErr := ioutil.ReadAll(res.Body)
	check(readErr)

	jsonErr := json.Unmarshal(body, &a)
	check(jsonErr)

	return a
}

func getVersionFromApi() Version {
	
	apiClient := http.Client {
		Timeout: time.Second * 2,
	}

	versionUrl := "http://" + c.ApiHost + ":" + c.ApiPort + "/version"
	versionReq, err := http.NewRequest(http.MethodGet, versionUrl, nil)
	check(err)

	versionRes, getErr := apiClient.Do(versionReq)
	check(getErr)

	body, readErr := ioutil.ReadAll(versionRes.Body)
	check(readErr)

	version := Version{}
	parseErr := json.Unmarshal(body, &version)
	check(parseErr)

	return version
}


func main() {
	log.Println("starting simmon web server...")
	log.Println("reading app config...")

	flag.Parse()
	if(!Filename.set) {
		log.Println("Configuration file not set (missing --conf <file> option). Trying default: /etc/simmon/simmon-web.conf")
		Filename.value = "/etc/simmon/simmon-web.conf"
	}

	doReadConfig(Filename.value)

	version := getVersionFromApi()

	log.Println("web portal v. " + c.Version)
	log.Println(" serving node/api version " + version.Version)

	tmpl := template.Must(template.ParseFiles("alerts.html"))
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData {
			Version: version.Version,
			Alerts: getAlertsFromApi(),
			Agents: getAgentsFromApi(),
		}

		tmpl.Execute(w, data)
	}).Methods("GET")

	log.Println("starting http listener on port :" + c.Port)
	http.ListenAndServe(":" + c.Port, router)
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
