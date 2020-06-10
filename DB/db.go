package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var dbCon *sql.DB

func main(){
	dbCon, _ := sql.Open("sqlite3", "./DB/example.sql")
	defer dbCon.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
				name STRING,
				age INT)`
	_, err := dbCon.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}


}
