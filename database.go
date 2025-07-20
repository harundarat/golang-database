package golangdatabase

import (
	"database/sql"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

var loadEnvOnce sync.Once

func GetConnection() *sql.DB {
	loadEnvOnce.Do(func() {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	})
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
