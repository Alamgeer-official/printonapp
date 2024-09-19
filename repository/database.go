package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var sqlDB *sql.DB
var gormDB *gorm.DB

// SqlDB returns the raw SQL database connection.
func SqlDB() *sql.DB {
	return sqlDB
}

// GormDB returns the GORM database connection.
func GormDB() *gorm.DB {
	return gormDB
}

// InitDbConnection initializes the database connection for both raw SQL and GORM.
func InitDbConnection() {
	// Load environment variables for database configuration
	loadEnvVars := func(key string) string {
		value := os.Getenv(key)
		if value == "" {
			log.Fatalf("Environment variable %s is not set", key)
		}
		return value
	}

	// Build the connection string for the SQL database
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		loadEnvVars("DB_HOST"),
		loadEnvVars("DB_PORT"),
		loadEnvVars("DB_USER"),
		loadEnvVars("DB_PWD"),
		loadEnvVars("DB_NAME"),
		loadEnvVars("DB_SETTING"),
	)

	// Open raw SQL database connection
	db, err := sql.Open(loadEnvVars("DB_DRIVER"), connectionString)
	if err != nil {
		log.Fatalf("Error connecting to SQL database: %v", err)
	}
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)
	sqlDB = db

	// Initialize GORM with the raw SQL database connection
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to GORM: %v", err)
	}

	// Set the search path for GORM
	if err := gormDB.Exec("SET search_path TO printonapp, public").Error; err != nil {
		log.Fatalf("Error setting search path: %v", err)
	}

	log.Println("Connected to database")
}
