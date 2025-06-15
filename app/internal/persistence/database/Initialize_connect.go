package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Initialize *gorm.DB
	sqlDB      *sql.DB
	closeOnce  sync.Once
)

func InitializeConnect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n[GORM] ", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Initialize = db

	sqlDBLocal, err := db.DB()
	if err != nil {
		log.Printf("Failed to retrieve sql.DB from GORM: %v", err)
		return
	}
	sqlDB = sqlDBLocal
}

func Close() {
	closeOnce.Do(func() {
		if sqlDB != nil {
			if err := sqlDB.Close(); err != nil {
				log.Printf("Failed to close database connection: %v", err)
			} else {
				log.Println("Database connection closed successfully.")
			}
		}
	})
}
