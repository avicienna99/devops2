package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Replace these values with your Azure MySQL Flexible Server details
	host := "testdb01.mysql.database.azure.com"
	port := 3306
	user := "nkouba@testdb01.mysql.database.azure.com"
	password := "Laithislockedin!"
	dbname := "reserveringsDB"

	// Connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Query the database
		rows, err := db.Query("SHOW TABLES")
		if err != nil {
			http.Error(w, "Failed to query the database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		fmt.Fprintln(w, "Tables in the database:")
		for rows.Next() {
			var tableName string
			err = rows.Scan(&tableName)
			if err != nil {
				http.Error(w, "Failed to scan row", http.StatusInternalServerError)
				return
			}
			fmt.Fprintln(w, tableName)
		}

		// Check for any error from iterating over rows
		if err = rows.Err(); err != nil {
			http.Error(w, "Error reading rows", http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
