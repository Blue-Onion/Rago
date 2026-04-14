package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	JWTSecert string
	Port      string
	DbUrl     string
}

func LoadConfig() *Config {
	godotenv.Load()
	dbUrl := os.Getenv("DBURL")
	port := os.Getenv("PORT")
	JWTSecert := os.Getenv("PORT")
	config := &Config{
		DbUrl:     dbUrl,
		Port:      port,
		JWTSecert: JWTSecert,
	}
	return config
}
