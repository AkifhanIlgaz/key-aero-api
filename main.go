package main

import (
	"log"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
	"github.com/AkifhanIlgaz/key-aero-api/db"
)

func main() {
	config, err := cfg.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not read environment variables", err)
	}

	databases, err := db.ConnectToDatabases(config)
	if err != nil {
		log.Fatal("Could not connect to databases: ", err)
	}

	defer databases.Postgres.Close()
	defer databases.Redis.Close()
}
