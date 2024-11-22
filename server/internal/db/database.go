package db

import (
	"fmt"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

var DB *gorm.DB
var logger = log.GetLogger()

func DbInit() {
		// PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	var err error
	// Retry loop for DB connection
	for {
		logger.Infof("Connecting to the database: %s", dsn)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Errorf("Database connection failed: %v. Retrying in 5 seconds... %v", err.Error())
			time.Sleep(5 * time.Second)
			continue
		} else {
			logger.Info("Successfully connected to the database.")
			break
		}
	}
}

func Migration(models ...any) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		logger.Error("Failed to migrate database: ", err)
	} else {
		logger.Info("Successfully migrated database.")
	}
}
