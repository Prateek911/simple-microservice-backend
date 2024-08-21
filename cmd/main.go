package main

import (
	"log"
	"simple-microservice-backend/api"
	"simple-microservice-backend/config"
	"simple-microservice-backend/db"
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

	envOpts, err := config.NewEnvConfig()
	if err != nil {
		log.Fatal("Error initiliasing Env")
	}

	if envOpts.Environment != "prod" && envOpts.Environment != "uat" {
		db.MigrateAndResetDB(db.DB)
	}

	server, err := api.NewServer(strconv.Itoa(opts.Host))
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	if err := server.Start(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
