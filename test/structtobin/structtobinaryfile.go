package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type ExampleData struct {
	Id       uint32
	Value    float32
	StrBytes [20]byte
}

type RiverLevelSensorData struct {
	Level       float32
	Temperature float32
	Pressure    float32
}

type RiverLevelApplicationData struct {
	UserId     uint64
	ApiKey     [16]byte
	SensorData []RiverLevelSensorData
}

func f1() {
	strBytes := [20]byte{}
	copy(strBytes[:], "bla")
	d := ExampleData{Id: 1, Value: 3.2, StrBytes: strBytes}
	exampleFilePath := "temp/data.bin"
	file, err := os.Create(exampleFilePath)
	if err != nil {
		fmt.Println("Error creating file ", err)
		return
	}
	defer file.Close()
	binary.Write(file, binary.LittleEndian, d)
	fmt.Println("Binary file created ", exampleFilePath)
}

func f2() {
	// create fixture data
	sensorData := []RiverLevelSensorData{
		{1, 30, 1030},
		{2, 33, 1030},
		{3, 35, 1030},
		{2, 33, 1030},
	}
	var apiKey = [16]byte{1, 2, 3, 4}
	var apiKeyBytes [16]byte
	copy(apiKeyBytes[:], apiKey[:])
	applicationData := RiverLevelApplicationData{3, apiKeyBytes, sensorData}
	// convert to raw bytes
	sendBytes := new(bytes.Buffer)
	binary.Write(sendBytes, binary.LittleEndian, applicationData.UserId)
	binary.Write(sendBytes, binary.LittleEndian, applicationData.ApiKey)
	binary.Write(sendBytes, binary.LittleEndian, uint32(len(applicationData.SensorData)))
	for _, v := range applicationData.SensorData {
		binary.Write(sendBytes, binary.LittleEndian, v)
	}
	fmt.Println("sendBytes ", sendBytes, " size ", sendBytes.Len())
	exampleFilePath := "temp/data.bin"
	file, err := os.Create(exampleFilePath)
	if err != nil {
		fmt.Println("Error creating file ", err)
		return
	}
	defer file.Close()
	file.Write(sendBytes.Bytes())
}

func main() {
	f2()
}
