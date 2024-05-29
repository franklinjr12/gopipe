package main

import (
	"encoding/binary"
	"fmt"
	"gopipe/internal/gopipeauth"
)

func testExtractUserAndKeyFromBytes() {
	var authData gopipeauth.DataInputAuth
	var headerBytes [24]byte
	authData.UserId = uint64(123)
	authData.ApiKey = "123"

	binary.LittleEndian.PutUint64(headerBytes[0:8], authData.UserId)
	copy(headerBytes[8:], []byte(authData.ApiKey))

	extractedData := gopipeauth.ExtractUserAndKeyFromBytes(headerBytes[:])

	fmt.Printf("expected UserId %v ApiKey %v\n", authData.UserId, authData.ApiKey)
	fmt.Printf("got UserId %v ApiKey %v\n", extractedData.UserId, extractedData.ApiKey)
}

func main() {
	testExtractUserAndKeyFromBytes()
}
