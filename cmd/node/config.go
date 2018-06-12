package main

/**
 * cmd/node/config.go
 * simmon - simple monitoring
 *
 * Config structure and functions.
 */

 import (
 	"encoding/json"
 	"flag"
 	"io/ioutil"
)

type Config struct {
	App			string 		`json:"app"`
	Listen		string 		`json:"listen"`
	Port		string 		`json:"port"`
	ApiListen	string		`json:"apiListen"`
	ApiPort		string		`json:"apiPort"`
	Storage		string 		`json:"storage"`
	DbUser		string 		`json:"dbuser"`
	DbPass		string 		`json:"dbpass"`
	DbHost		string 		`json:"dbhost"`
	DbPort		string 		`json:"dbport"`
	DbName		string		`json:"dbname"`
	DbTable		string 		`json:"dbtable"`
	Version		string		`json:"version"`
}

type confFlag struct {
	set bool
	value string
}

var Filename confFlag

func init() {
	flag.Var(&Filename, "conf", "simmon node configuration")
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