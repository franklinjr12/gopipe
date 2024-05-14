package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Open() (db *sql.DB) {
	// obviously this should be read from environment variable or somewhere safe
	const DSN = "host=localhost port=5432 user=postgres password=postgres dbname=gopipe sslmode=disable"
	const DATABASE_DRIVER = "postgres"
	db, err := sql.Open(DATABASE_DRIVER, DSN)
	if err != nil {
		fmt.Println("Error on sql.Open")
		return nil
	}
	// ping to confirm connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error on db.Ping")
		return nil
	}
	return db
}

func UserExist(db *sql.DB, userId uint64, apiKey string) bool {
	var exist bool
	err := db.QueryRow("select exists (select 1 from users where id = $1 and apikey = $2)", userId, apiKey).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}
