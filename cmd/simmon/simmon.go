package main

/**
 * simmon.go
 * simmon - simple monitoring
 *
 * CLI utilities for managing the local monitoring node.
 */

import (
 	"encoding/json"
 //	"fmt"
 	"io/ioutil"
)

type Config struct {
	App			string 		`json:"app"`
	Listen		string 		`json:"listen"`
	Port		string 		`json:"port"`
	Storage		string 		`json:"storage"`
	DbUser		string 		`json:"dbuser"`
	DbPass		string 		`json:"dbpass"`
	DbHost		string 		`json:"dbhost"`
	DbPort		string 		`json:"dbport"`
	DbTable		string 		`json:"dbtable"`
}

var c Config

func doReadConfig(cfile string) {
	cfg, _ := ioutil.ReadFile(cfile)
	
	err := json.Unmarshal(cfg, &c)
	check(err)
}

func main() {

	doReadConfig("../../configs/node/simmon-node.conf")

	simchan := make(chan string, 100)


	simchan <- "GET_ALERTS"

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}