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

func main() {
	testSelectApplicationId()
	// testUserExist()
}
