package dataingestion

import (
	"database/sql"
	"fmt"
)

type DataIngestionInput struct {
	Data        []byte
	Application string
	UserId      uint64
	Args        []string
}

func Ingest(dataInput DataIngestionInput) {
	db, err := sql.Open("postgres", "gopipePostgres")
	if err != nil {
		fmt.Println("check database setup")
		return
	}
	defer db.Close()
}
