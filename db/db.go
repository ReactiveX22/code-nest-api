package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	DB   *bun.DB
	once sync.Once
)

func InitDB() error {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	once.Do(func() {
		dsn := getDSN()

		sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

		DB = bun.NewDB(sqlDB, pgdialect.New())

		if err = DB.PingContext(context.Background()); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		log.Println("Database successfully connected!")
	})

	return err
}

func getDSN() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname)
}
