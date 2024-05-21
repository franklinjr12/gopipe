package dataingestion

import (
	"fmt"
	"gopipe/internal/database"
	"gopipe/internal/normalizer"
)

type DataIngestionInput struct {
	Data        []byte
	Application string
	UserId      uint64
	DeviceId    uint64
	ShouldPipe  string
	Args        []string
}

func Ingest(dataInput DataIngestionInput) {
	db := database.Open()
	if db == nil {
		return
	}
	defer db.Close()
	fmt.Println("Doing database work...")
	if dataInput.ShouldPipe == "" || dataInput.ShouldPipe == "false" {
		//imediate store
		err := database.WriteUnpipedData(db, database.UnpipedData{UserId: dataInput.UserId, Application: dataInput.Application, Data: dataInput.Data})
		if err != nil {
			fmt.Println("Error writting unpiped data. ", err.Error())
		}
	} else {
		jsonData, err := normalizer.ToJson(dataInput.Data)
		if err != nil {
			// data is in raw bytes format, need to parse it to structs
			fmt.Println("data is in raw bytes format, need to parse it to structs")
		} else {
			fmt.Println("Ignore. ", jsonData)
		}
	}
}
