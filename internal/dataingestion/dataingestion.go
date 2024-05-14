package dataingestion

import (
	"fmt"
	"gopipe/internal/database"
)

type DataIngestionInput struct {
	Data        []byte
	Application string
	UserId      uint64
	Args        []string
}

func Ingest(dataInput DataIngestionInput) {
	db := database.Open()
	if db == nil {
		return
	}
	defer db.Close()
	fmt.Println("Doing database work...")
}
