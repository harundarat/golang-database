package golangdatabase

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", os.Getenv("DB_SOURCENAME"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// use db here
}
