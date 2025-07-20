package golangdatabase

import (
	"context"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestExecSql(t *testing.T) {
	envErr := godotenv.Load()
	if envErr != nil {
		panic(envErr)
	}
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES('harun', 'Harun Al Rasyid')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert into customer")
}
