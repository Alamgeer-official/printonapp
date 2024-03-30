package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var sqlDB *sql.DB

func SqlDB() *sql.DB {
	return sqlDB
}

var gormDB *gorm.DB

func GormDB() *gorm.DB {
	return gormDB
}

func InitDbConnectionos() {
	//load db config file from env
	env := os.Getenv
	_ = env
	// Connection opening string
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", env("DB_HOST"), env("DB_PORT"), env("DB_USER"), env("DB_PWD"), env("DB_NAME"), env("DB_SETTING"))

	//connecting to DB
	db, err := sql.Open(env("DB_DRIVER"), connectionString)
	if err != nil {
		log.Fatalf("unable to connect db : %v", err)
	}
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)
	sqlDB = db

	// Connect Orm
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect gorm : %v", err)
	}
	gormDB = gormdb
	log.Println("connected to db")
}
