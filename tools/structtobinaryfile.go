package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type ExampleData struct {
	Id       uint32
	Value    float32
	StrBytes [20]byte
}

func main() {
	strBytes := [20]byte{}
	copy(strBytes[:], "bla")
	d := ExampleData{Id: 1, Value: 3.2, StrBytes: strBytes}
	file, err := os.Create("temp/data.bin")
	if err != nil {
		fmt.Println("Error creating file ", err)
		return
	}
	defer file.Close()
	binary.Write(file, binary.LittleEndian, d)
	fmt.Println("Binary file created")
}
