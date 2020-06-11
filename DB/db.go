package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var dbCon *sql.DB

type Person struct {
	Name string
	Age int
}

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

//	singleSelect(dbCon)
//	multipleSelect(dbCon)
//	insert(dbCon)
//	update(dbCon)
//	delete(dbCon)
	sqlInjection(dbCon)


}

func update(dbCon *sql.DB) {
	updateCmd := "UPDATE person SET age = ? WHERE name = ?"
	_, err := dbCon.Exec(updateCmd,80, "Tank")

	if err != nil {
		log.Fatalln(err)
	}

}

func insert(dbCon *sql.DB) {
	insertCmd := "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err := dbCon.Exec(insertCmd, "Tank", 29)

	if err != nil {
		log.Fatalln(err)
	}
}
func delete(dbCon *sql.DB) {
	deleteCmd := "DELETE FROM person WHERE name = ?"
	_, err := dbCon.Exec(deleteCmd, "Nancy")

	if err != nil {
		log.Fatalln(err)
	}
}

func multipleSelect(dbCon *sql.DB) {
	selectCmd := "SELECT * FROM person"
	rows, _ := dbCon.Query(selectCmd)
	fmt.Println("rows: ", rows)
	defer rows.Close()
	var persons []Person
	for rows.Next() {
		var p Person
		rows.Scan(&p.Name, &p.Age)

		persons = append(persons, p)
	}
	err := rows.Err()
	if err != nil {
		log.Println()
	}
	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}

}

func singleSelect(dbCon *sql.DB) {
	selectCmd := "SELECT * FROM person WHERE age = ?"
	row := dbCon.QueryRow(selectCmd,200)
	var p Person
	err := row.Scan(&p.Name, &p.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("NORows")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

}

func sqlInjection(dbCon *sql.DB) {
	tableName := "person; INSERT INTO person (name, age) VALUES ('inject',100)"
	selectCmd := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := dbCon.Query(selectCmd)
	fmt.Println("rows: ", rows)
	defer rows.Close()
	var persons []Person
	for rows.Next() {
		var p Person
		rows.Scan(&p.Name, &p.Age)

		persons = append(persons, p)
	}
	err := rows.Err()
	if err != nil {
		log.Println()
	}
	for _, p := range persons {
		fmt.Println(p.Name, p.Age)
	}

}