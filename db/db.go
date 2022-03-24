package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	PORT = 5432
// )
// ErrNoMatch is returned when request a row that doesn't exist
// ErrNoMatch dikembalikan ketika request sebuah row tidak ada / tidak tersedia

var ErrNoMatch = fmt.Errorf("no matching record")

var dB *gorm.DB

func DatabaseInitialize() *gorm.DB {
	var err error

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASS")
	// Initialize DB
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			host, port, dbName, username, password),
	), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("[INIT] Failed connecting to PostgreSQL Database at %s:%s. %+v\n",
			host, port, err)
	}
	log.Printf("[INIT] Successfully connected to PostgreSQL Database\n")
	sqlDb, err := db.DB()
	sqlDb.SetMaxOpenConns(100)
	dB = db
	return dB
}
