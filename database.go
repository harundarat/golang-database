package golangdatabase

import (
	"database/sql"
	"os"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_SOURCENAME"))
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db
}
