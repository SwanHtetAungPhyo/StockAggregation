package db

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
var logger = log.GetLogger()

func DbInit() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dsn := "root:NewPassword@tcp(127.0.0.1:3306)/stockAgg?charset=utf8mb4&parseTime=True&loc=Local"

	logger.Infof("Connecting to the database: %s", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Database connection failed: %v. Retrying in 5 seconds... %v", err.Error())
		time.Sleep(5 * time.Second)
	} else {
		logger.Info("Successfully connected to the database.")
	}

}

func Migration(models ...any) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		logger.Error("Failed to migrate database")
		return
	}
}
