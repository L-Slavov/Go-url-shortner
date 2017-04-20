package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/ziutek/mymysql/godrv"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "urlshortner"
	DB_USER = "root"
	DB_PASS = ""
)

func OpenDB() *sql.DB {

	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	return db
}

func DBInsert(url string) int64 {
	db := OpenDB()
	defer db.Close()
	result, err := db.Exec("INSERT INTO urls (url) VALUES (?)", url)
	if err != nil {
		fmt.Println("Error: couldn't insert url")
		fmt.Println(err)
	}
	returnValue, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Couldn't get lastIndex")
	}
	return returnValue
}

func DBget(id int) string {
	db := OpenDB()
	defer db.Close()
	var link string
	err := db.QueryRow("SELECT url FROM urls WHERE id = ?", id).Scan(&link)
	if err != nil {
		fmt.Println("Couldn't get link from database")
		fmt.Println(err)
	}

	return link
}
