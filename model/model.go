package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func init() {
	db, err := sql.Open("postgres", "user=postgres password=admin dbname=cloudpiece sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	conn = db
}
