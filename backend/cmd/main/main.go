package main

import (
	"fmt"
	"phone-accounting-system/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	
	DBConfig := database.Configuration{
		Host: "localhost",
		Port: "6001",
		User: "admin",
		Password: "1234",
		DBName: "phone_accounting_db",
	}

	db, err := database.NewDB(&DBConfig)

	if err != nil {
		fmt.Println("Error: %w", err)
		return
	}

	if db == nil {
		fmt.Println("DB is NULL")
		return
	}

	var result string 

	err = db.QueryRow("Select login from \"user\" LIMIT 1").Scan(&result)

	if err != nil {
		fmt.Println("Select error: %w", err)
		return
	}

	fmt.Println(result)

}