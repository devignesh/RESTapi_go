package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	database, err := sql.Open("mysql", "root:root(127.0.0.1:3306)/gomysql")

	if err != nil {
		panic(err.Error())
	}

	defer database.Close()

	fmt.Println("db connected")

	insert, err := database.Query("INSERT INTO test VALUES('vicky)")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("data inserted successfully")
}
