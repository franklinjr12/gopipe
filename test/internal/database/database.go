package main

import (
	"fmt"
	"gopipe/internal/database"
)

func testSelectApplicationId() {
	db := database.Open()
	var companyId uint64
	queryStr := "select id from companies where name = $1"
	err := db.QueryRow(queryStr, "River Safety CO").Scan(&companyId)
	if err != nil {
		fmt.Println("Error reading company_id: ", err)
		return
	}
	fmt.Println("Company: ", companyId)
	applicationId := database.SelectApplicationId(db, "River level monitoring", companyId)
	fmt.Println("ApplicationId ", applicationId)
}

func testUserExist() {
	db := database.Open()
	exist := database.UserExist(db, 3, "1234567890123456")
	fmt.Println("Exist: ", exist)
}

func testSelectApplicationDataStructure() {
	db := database.Open()
	applicationId := uint64(1)
	version := 0
	rows := database.SelectApplicationDataStructure(db, applicationId, version)
	if rows != nil {
		defer rows.Close()
		columns, _ := rows.Columns()
		var dbTypeNames []string
		types, _ := rows.ColumnTypes()
		for _, v := range types {
			dbTypeNames = append(dbTypeNames, v.DatabaseTypeName())
		}
		fmt.Printf("Columns:\t%v\nTypes:\t\t%v\n", columns, dbTypeNames)
	}
}

func testInsertPipedData() {
	pipedData := database.PipedData{
		ApplicationId: 1, Version: 0, Data: []byte{},
	}
	db := database.Open()
	err := database.InsertPipedData(db, &pipedData)
	fmt.Println("InsertPipedData err: ", err)
}

func main() {
	// testUserExist()
	// testSelectApplicationId()
	// testSelectApplicationDataStructure()
	testInsertPipedData()
}
