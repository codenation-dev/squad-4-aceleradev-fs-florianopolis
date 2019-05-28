package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("DB connected!!!")
	}
}

func createDB(db *DB) {
	_, err := db.Exec("CREATE DATABASE testDB")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}
}
