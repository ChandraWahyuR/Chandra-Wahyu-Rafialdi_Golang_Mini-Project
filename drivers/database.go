package drivers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBSsl  string
}

func ConnectDB(con Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		con.DBHost,
		con.DBUser,
		con.DBPass,
		con.DBName,
		con.DBPort,
		con.DBSsl,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := createDatabase(db, con.DBName); err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	MigrationUser(db)
	return db
}

func MigrationUser(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Equipment{}, &Rent{})
}

func createDatabase(db *gorm.DB, dbname string) error {
	var count int64
	if err := db.Raw("SELECT COUNT(*) FROM pg_database WHERE datname = ?", dbname).Scan(&count).Error; err != nil {
		return err
	}

	// Make db
	if count == 0 {
		if err := db.Exec("CREATE DATABASE " + dbname).Error; err != nil {
			return err
		}
	}

	return nil
}
