package config

import (
	"os"
	// "log"
	// "fmt"
	//"github.com/joho/godotenv"
)

type Config struct {
	PostgreDriver  string "POSTGRES_DRIVER"
	Host string "POSTGRES_HOST"
	PortPG string "POSTGRES_PORT"
	User string "POSTGRES_USER"
	DBName string "POSTGRES_DB"
	Password string "POSTGRES_PASSWORD"
	MigrationsSource string "MIGRATIONS_SOURCE"
	Port string "PORT"
}

func LoadConfig(path string) (config Config, err error) {
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error getting env, %v", err)
	// } else {
	// 	fmt.Println("We are getting values")
	// }
	return Config{
		PostgreDriver:  os.Getenv("POSTGRES_DRIVER"),
		Host: os.Getenv("POSTGRES_HOST"),
		PortPG: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		DBName: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		MigrationsSource: os.Getenv("MIGRATIONS_SOURCE"),
		Port: os.Getenv("PORT"),
	}, err
}
