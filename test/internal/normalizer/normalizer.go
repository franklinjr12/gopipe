package main

import (
	"fmt"
	"gopipe/internal/database"
	"gopipe/internal/normalizer"
)

func testRowsToFormatStruct() {
	db := database.Open()
	rows := database.SelectApplicationDataStructure(db, 1, 0)
	appStruct := normalizer.RowsToFormatStruct(rows)
	fmt.Println("appStruct: ", appStruct)
}

func main() {
	testRowsToFormatStruct()
}
