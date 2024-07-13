package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)

}
