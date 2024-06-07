package main

import (
	"bytes"
	"encoding/binary"
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

func testBytesToStruct() {
	format := []normalizer.ApplicationBytesDecode{
		{FirstByte: 0, LastByte: 3, Type: "int"},
		{FirstByte: 4, LastByte: 6, Type: "bytes"},
	}
	id := int32(1)
	str := "123"
	dataBytes := new(bytes.Buffer)
	binary.Write(dataBytes, binary.LittleEndian, id)
	for i, _ := range str {
		binary.Write(dataBytes, binary.LittleEndian, str[i])
	}
	fmt.Printf("DataBytes: %v len: %v\n", dataBytes.Bytes(), len(dataBytes.Bytes()))
	result := normalizer.BytesToStruct(dataBytes.Bytes(), 1, format)
	fmt.Println("Result: ", result)
}

func main() {
	testRowsToFormatStruct()
	testBytesToStruct()
}
