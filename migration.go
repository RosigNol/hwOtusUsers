package main

import (
	"log"
	"otusHWUsers/config"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/google/uuid"
)


func main() {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	
	psqlInfo := "postgres://"+ config.User + ":"+ config.Password + "@"+ config.Host +":" + config.PortPG+ "/" + config.DBName + "?sslmode=disable"

	m, err := migrate.New(
		config.MigrationsSource,
		psqlInfo)
	if err != nil {
		log.Fatalf("connection fail: %v", err, psqlInfo)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("migration fail: %v", err)
	}

}
