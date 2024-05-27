package main

import (
	"encoding/binary"
	"fmt"
	"gopipe/internal/normalizer"
	"math"
	"os"
)

func f1() {
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
}

func f2() {
	riverDataFormat := []normalizer.ApplicationBytesDecode{
		{0, 4, "float"},
		{4, 8, "float"},
		{8, 12, "float"},
	}
	riverDataPacketSize := 12
	// read binary data
	file, err := os.Open("temp/data.bin")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()
	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info ", err)
		return
	}
	fileSize := fileInfo.Size()
	fmt.Println("Filesize ", fileSize)
	var fileData = make([]byte, fileSize)
	err = binary.Read(file, binary.LittleEndian, &fileData)
	if err != nil {
		fmt.Println("Error reading file ", err)
		return
	}
	//first 24 bytes are header containing 8 bytes id and 16 bytes apikey
	userId := uint64(binary.LittleEndian.Uint64(fileData[0:8]))
	apiKey := fileData[8:24]
	dataBytes := fileData[28:]
	numPackets := uint32(binary.LittleEndian.Uint32(fileData[24:28]))
	fmt.Println("dataBytes: ", dataBytes)
	fmt.Println("dataBytes lenght ", len(dataBytes))
	fmt.Println("numPackets ", numPackets)
	var applicationData []any
	for i := 0; i < int(numPackets); i++ {
		packet := dataBytes[i*riverDataPacketSize : i*riverDataPacketSize+riverDataPacketSize]
		for _, v := range riverDataFormat {
			switch v.Type {
			case "int":
				applicationData = append(applicationData, int32(binary.LittleEndian.Uint32(packet[v.FirstByte:v.LastByte])))
			case "float":
				applicationData = append(applicationData, math.Float32frombits(binary.LittleEndian.Uint32(packet[v.FirstByte:v.LastByte])))
			case "bytes":
				applicationData = append(applicationData, string(packet[v.FirstByte:v.LastByte]))
			}
		}
	}
	fmt.Println("User ", userId, " key ", apiKey, " data ", applicationData)
}

func main() {
	// f1()
	f2()
}
