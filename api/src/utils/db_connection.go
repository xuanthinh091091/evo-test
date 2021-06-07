package utils

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// CreateConnection for api
func CreateConnection() *sql.DB {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		dburl = "postgres://postgres:postgres@localhost:5432/book_db?sslmode=disable"
	}
	db, err := sql.Open("postgres", dburl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
