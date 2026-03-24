package main

import (
	"fmt"
	"phone-accounting-system/internal/database"
	"phone-accounting-system/internal/repository"

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

	repo := repository.UserRepo{DB: db.DB}

	arr := repo.GetAllUsers()

	for _, v := range arr {
		fmt.Println(v)
	}

}