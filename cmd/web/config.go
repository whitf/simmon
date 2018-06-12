package main

/**
 * cmd/web/config.go
 * simmon - simple monitoring
 *
 * Config structures and functions for the simmon web portal.
 */


import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type Config struct {
	App					string		`json:"app"`
	ApiHost				string		`json:"apiHost"`
	ApiPort				string		`json:"apiPort"`
	Listen				string		`json:"listen"`
	Port 				string		`json:"port"`
	Version				string		`json:"version"`
}

type confFlag struct {
	set bool
	value string
}

var Filename confFlag

func init() {
	flag.Var(&Filename, "conf", "simmon web server configuration")
}

//@TODO: add in better error checking (file not found, etc)
func doReadConfig(cfile string) {
	cfg, _ := ioutil.ReadFile(cfile)
	
	err := json.Unmarshal(cfg, &c)
	check(err)

}

func (cf *confFlag) Set(x string) error {
	cf.value = x
	cf.set = true
	return nil
}

func (cf *confFlag) String() string {
	return cf.value
}