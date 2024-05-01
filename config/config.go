package config

import (
	"log"
	"os"

	db "prototype/drivers"

	"github.com/joho/godotenv"
)

func LoadFileEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Eror, while loading .env file")
	}
}

func InitConfigDb() db.Config {
	return db.Config{
		DBName: os.Getenv("DBName"),
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBHost: os.Getenv("DBHost"),
		DBPort: os.Getenv("DBPort"),
		DBSsl:  os.Getenv("DBSsl"),
	}
}
