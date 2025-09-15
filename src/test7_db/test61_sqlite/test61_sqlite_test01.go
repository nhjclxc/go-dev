package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"unicode/utf8"
)

func main1() {

	data := "♥"
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS anonymous_user (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO anonymous_user(name) VALUES (?)", "Alice")
	if err != nil {
		log.Fatal(err)
	}
}
