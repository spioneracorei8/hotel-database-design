package main

import (
	"hotel-database-design/helper"
	"hotel-database-design/logger"
	"hotel-database-design/models"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RUN MIGRATE TABLES
/*
go run "/home/spionera/Documents/me/code/hotel-database-design/migration/migration.go"
*/
func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error loading .env file: %v", err)
	}
	gormLogger := logger.GormLogger()
	PSQL_CONNECTION := helper.GetENV("PSQL_CONNECTION", "")
	db, err := gorm.Open(postgres.Open(PSQL_CONNECTION), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		logrus.Errorf("Error connecting to database: %v", err)
	}
	if err := db.Migrator().AutoMigrate(
		&models.Hotel{},
	); err != nil {
		logrus.Errorf("Error migrating database table: %v", err)
	}

}
