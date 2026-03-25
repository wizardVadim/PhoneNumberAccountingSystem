package main

import (
	"fmt"
	"log"
	"net/http"
	"phone-accounting-system/internal/database"
	"phone-accounting-system/internal/handlers"
	"phone-accounting-system/internal/repository"

	_ "github.com/lib/pq"
)

func main() {

	DBConfig := database.Configuration{
		Host:     "localhost",
		Port:     "6001",
		User:     "admin",
		Password: "1234",
		DBName:   "phone_accounting_db",
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

	defer db.Close()

	userRepo := &repository.UserRepo{DB: db.DB}
	roleRepo := &repository.UserRoleRepo{DB: db.DB}
	personRepo := &repository.PhysicalPersonRepo{DB: db.DB}
	phoneRepo := &repository.PhoneNumberRepo{DB: db.DB}
	phoneTypeRepo := &repository.PhoneNumberTypeRepo{DB: db.DB}

	mux := http.NewServeMux()

	ah := handlers.AuthHandler{UserRepo: userRepo, Mux: mux}
	ph := handlers.PersonHandler{PersonRepo: personRepo, Mux: mux}
	phh := handlers.PhoneHandler{PhoneRepo: phoneRepo, PhoneTypeRepo: phoneTypeRepo, Mux: mux}
	uh := handlers.UserHandler{UserRepo: userRepo, UserRoleRepo: roleRepo, Mux: mux}

	ah.Init()
	ph.Init()
	phh.Init()
	uh.Init()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
