package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func initStorage() *sql.DB {
	_, err := os.Stat("coupon.sqlite3")

	if os.IsNotExist(err) {
		log.Print("no db file")
		os.Create("coupon.sqlite3")
	}

	db, err := sql.Open("sqlite3", "./coupon.db")
	if err != nil {
		panic(err)
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS coupon (id INTEGER PRIMARY KEY, name TEXT, brand TEXT, value INT, createdAt DATETIME, expiry DATETIME)")
	if err != nil {
		panic(err)
	}
	statement.Exec()

	// statement, err = db.Prepare("INSERT INTO coupon (name, brand, value, createdAt) VALUES (?, ?)")
	// statement.Exec("Tesco 20 off", "Tesco", 20, time.Now())
	// if err != nil {
	// 	panic(err)
	// }

	return db
}
