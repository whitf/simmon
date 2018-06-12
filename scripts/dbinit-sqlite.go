package main

/**
 * dbinit-sqlite.go
 */


import (
 	"database/sql"
 	"encoding/json"
 	"fmt"
 	"io/ioutil"

 	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	App					string 		`json:"app"`
	Listen				string 		`json:"listen"`
	Port				string 		`json:"port"`
	Storage				string 		`json:"storage"`
	DbUser				string 		`json:"dbuser"`
	DbPass				string 		`json:"dbpass"`
	DbHost				string 		`json:"dbhost"`
	DbPort				string 		`json:"dbport"`
	DbName				string 		`json:"dbname"`
	DbArchiveFreq		string 		`json:"db_archive_frequency"`
}

var c Config

const (
	CREATE_QOS_TABLE = "CREATE TABLE IF NOT EXISTS qos(alert BOOLEAN, alertLevel TEXT, alertValue TEXT, agent TEXT, agentAddress TEXT, agentId TEXT, agentVersion TEXT, heartbeatThreshold FLOAT, date TEXT, msg TEXT, qosType TEXT)"
)

func doReadConfig(cfile string) {
	cfg, _ := ioutil.ReadFile(cfile)

	err := json.Unmarshal(cfg, &c)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("creating local sqlite database")
	fmt.Println()

	fmt.Println("reading config...")
	doReadConfig("/etc/simmon/simmon-node.conf")
	fmt.Println(" done")

	fmt.Println("connecting to database...")
	database, databaseErr := sql.Open("sqlite3", c.DbName)
	if databaseErr != nil {
		fmt.Println(" error connecting to and/or creating database")
	} else {
		fmt.Println(" success")
	}
	fmt.Println(" done")

	fmt.Println("creating QoS table...")
	statement, _ := database.Prepare(CREATE_QOS_TABLE)
	statement.Exec()
	fmt.Println(" done")

}

