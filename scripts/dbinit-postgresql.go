package main

/**
 * dbinit-postgresql.go
 * database init for postgres
 */

import (
 	"database/sql"
 	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "github.com/lib/pq"
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
	CREATE_UUID_GENERATOR = "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;"
	CREATE_AGENT_TABLE = "CREATE TABLE agent(id uuid, alive boolean, address varchar(128), name varchar(256));"
	CREATE_QOS_TABLE = "CREATE TABLE qos(agentid varchar(128), alert boolean, alert_level varchar(64), alert_value varchar(256), agent varchar(256), msg varchar(256), type varchar(64), id uuid NOT NULL DEFAULT uuid_generate_v1(), sdate date NOT NULL DEFAULT CURRENT_DATE);"
	CREATE_ALERT_TABLE = "CREATE TABLE alert(id uuid, qosid uuid, agentid uuid, thrown varchar(128));"
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
	fmt.Println("starting postgres db init")
	fmt.Println()
	fmt.Println("creating new simmon table structure in postgresql")
	fmt.Println("based on the db config in ../configs/node/simmon-node.conf")
	fmt.Println()

	doReadConfig("../configs/node/simmon-node.conf")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		c.DbUser, c.DbPass, c.DbName, c.DbHost)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	//stmt, dberr := db.Prepare(CREATE_QOS_TABLE)
	//checkErr(dberr)

	fmt.Println("running")

	fmt.Println(" CREATE UUID generator (if necessary)")
	fmt.Println("  " + CREATE_UUID_GENERATOR)

	uuidStmt, uuidErr := db.Prepare(CREATE_UUID_GENERATOR)
	checkErr(uuidErr)
	_, err = uuidStmt.Exec()
	checkErr(err)

	fmt.Println("   done")

	fmt.Println(" CREATE TABLE agent")
	fmt.Println("  " + CREATE_AGENT_TABLE)

	stmt, createErr := db.Prepare(CREATE_AGENT_TABLE)
	checkErr(createErr)
	_, err = stmt.Exec()
	checkErr(err)

	fmt.Println("   done")

	fmt.Println(" CREATE TABLE qos")
	fmt.Println("  " + CREATE_QOS_TABLE)

	stmt, createErr = db.Prepare(CREATE_QOS_TABLE)
	checkErr(createErr)
	_, err = stmt.Exec()
	checkErr(err)

	fmt.Println("   done")
	
	fmt.Println(" CREATE TABLE alert")
	fmt.Println("  " + CREATE_ALERT_TABLE)

	stmt, createErr = db.Prepare(CREATE_ALERT_TABLE)
	checkErr(createErr)
	_, err = stmt.Exec()
	checkErr(err)

	fmt.Println("   done")


	//_, err = stmt.Exec()
	//checkErr(err)

	fmt.Println()
	fmt.Println("done")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
