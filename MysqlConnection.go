package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := GetDbConnection()

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}

func GetDbConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:suhumar123@tcp(127.0.0.1:3306)/websitedata")

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Ping failed")
		panic(pingErr.Error())
	}

	// if there is an error opening the connection, handle it
	if err != nil {

		fmt.Println("error")
		panic(err.Error())
	}

	if err == nil {
		fmt.Println("Making a connection with mysql")
	}

	return db

}
