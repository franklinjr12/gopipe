package main

import (
	"encoding/binary"
	"fmt"
	"gopipe/internal/normalizer"
	"math"
	"os"
)

func main() {
	exampleApplicationDecodeFormat := []normalizer.ApplicationBytesDecode{
		{0, 4, "int"},
		{4, 8, "float"},
		{8, 28, "bytes"},
	}
	// read binary data
	file, err := os.Open("temp/data.bin")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()
	var fileData [28]byte
	err = binary.Read(file, binary.LittleEndian, &fileData)
	if err != nil {
		fmt.Println("Error reading file ", err)
		return
	}
	var applicationData []any
	for _, e := range exampleApplicationDecodeFormat {
		fmt.Printf("Converting %d %d to %s\n", e.FirstByte, e.LastByte, e.Type)
		switch e.Type {
		case "int":
			applicationData = append(applicationData, int32(binary.LittleEndian.Uint32(fileData[e.FirstByte:e.LastByte])))
		case "float":
			applicationData = append(applicationData, math.Float32frombits(binary.LittleEndian.Uint32(fileData[e.FirstByte:e.LastByte])))
		case "bytes":
			applicationData = append(applicationData, string(fileData[e.FirstByte:e.LastByte]))
		}
	}
	fmt.Println("Conversion done: ", applicationData)
	// outputs -> Conversion done:  [1 0]
	// should output -> Conversion done:  [1 3.2 "bla"]
}
