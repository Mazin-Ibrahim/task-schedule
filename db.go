package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	connStr := "user=mazin dbname=task sslmode=disable password=password host=localhost port=5432"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func closeDB(db *sql.DB) {

	db.Close()
}
