package main

import (
	"log"
	"simple-microservice-backend/api"
	"simple-microservice-backend/config"
	"simple-microservice-backend/db"
	"simple-microservice-backend/db/model"
	"strconv"
)

func init() {
	db.InitDB()
}

func main() {
	opts, err := config.NewServerConfig()
	if err != nil {
		log.Fatal("Error initiliasing server")
	}

	models := []interface{}{&model.AccountMaster{},
		&model.Employee{},
		&model.Owner{},
		&model.Payments{},
		&model.Contact{},
		&model.Contactables{}}

	for _, model := range models {
		if err := db.DB.AutoMigrate(model); err != nil {
			log.Fatalf("Error migrating %T, %v", model, err)
		}

	}

	server, err := api.NewServer(strconv.Itoa(opts.Host))
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
