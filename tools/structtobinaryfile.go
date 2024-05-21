package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type ExampleData struct {
	id       uint64
	value    float32
	strBytes [20]byte
}

func main() {
	strBytes := [20]byte{}
	copy(strBytes[:], "bla")
	d := ExampleData{id: 1, value: 3.2, strBytes: strBytes}
	file, err := os.Create("temp/data.bin")
	if err != nil {
		fmt.Println("Error creating file ", err)
		return
	}
	defer file.Close()
	binary.Write(file, binary.LittleEndian, d)
	fmt.Println("Binary file created")
}
