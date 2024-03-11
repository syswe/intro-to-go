package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "postgres://username:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping the database:", err)
		return
	}

	// Execute a query
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			return
		}
		fmt.Println("ID:", id, "Name:", name)
	}

	// Check for any errors during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during iteration:", err)
		return
	}
}
