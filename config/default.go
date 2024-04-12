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
	Origin string "ORIGIN"

	AccessTokenPrivateKey  string        "ACCESS_TOKEN_PRIVATE_KEY"
	AccessTokenPublicKey   string        "ACCESS_TOKEN_PUBLIC_KEY"
	RefreshTokenPrivateKey string        "REFRESH_TOKEN_PRIVATE_KEY"
	RefreshTokenPublicKey  string        "REFRESH_TOKEN_PUBLIC_KEY"
	AccessTokenExpiresIn   time.Duration "ACCESS_TOKEN_EXPIRED_IN"
	RefreshTokenExpiresIn  time.Duration "REFRESH_TOKEN_EXPIRED_IN"
	AccessTokenMaxAge      int           "ACCESS_TOKEN_MAXAGE"
	RefreshTokenMaxAge     int           "REFRESH_TOKEN_MAXAGE"
}

func LoadConfig(path string) (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values")
	}
	return Config{
		PostgreDriver:  os.Getenv("POSTGRES_DRIVER"),
		Host: os.Getenv("POSTGRES_HOST"),
		PortPG: os.Getenv("POSTGRES_PORT"),
		User: os.Getenv("POSTGRES_USER"),
		DBName: os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		MigrationsSource: os.Getenv("MIGRATIONS_SOURCE"),
		Origin: os.Getenv("ORIGIN"),
		AccessTokenPrivateKey: os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"),
		AccessTokenPublicKey: os.Getenv("ACCESS_TOKEN_PUBLIC_KEY"),
		RefreshTokenPrivateKey: os.Getenv("REFRESH_TOKEN_PRIVATE_KEY"),
		RefreshTokenPublicKey: os.Getenv("REFRESH_TOKEN_PUBLIC_KEY"),
		AccessTokenExpiresIn: os.Getenv("ACCESS_TOKEN_EXPIRED_IN"),
		RefreshTokenExpiresIn: os.Getenv("REFRESH_TOKEN_EXPIRED_IN"),
		AccessTokenMaxAge: os.Getenv("ACCESS_TOKEN_MAXAGE"),
		RefreshTokenMaxAge: os.Getenv("REFRESH_TOKEN_MAXAGE"),
	}, err
}
