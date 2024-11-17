package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	// Replace these values with your Azure MySQL Flexible Server details
	host := "testdb01.mysql.database.azure.com"
	port := 3306
	user := "nkouba@testdb01.mysql.database.azure.com"
	password := "Laithislockedin!"
	dbname := "reserveringsDB"

	// Connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=true", user, password, host, port, dbname)

	// Open the database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Query the database (example: list all tables)
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("Failed to query the database: %v", err)
	}
	defer rows.Close()

	fmt.Println("Tables in the database:")
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			log.Fatalf("Failed to scan row: %v", err)
		}
		fmt.Println(tableName)
	}

	// Check for any error from iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatalf("Error reading rows: %v", err)
	}
}
