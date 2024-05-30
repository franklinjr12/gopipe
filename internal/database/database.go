package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type UnpipedData struct {
	Data        []byte
	Application string
	UserId      uint64
}

func Open() (db *sql.DB) {
	// obviously this should be read from environment variable or somewhere safe
	const DSN = "host=localhost port=5432 user=postgres password=postgres dbname=gopipe sslmode=disable"
	const DATABASE_DRIVER = "postgres"
	db, err := sql.Open(DATABASE_DRIVER, DSN)
	if err != nil {
		fmt.Println("Error on sql.Open")
		return nil
	}
	// ping to confirm connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error on db.Ping")
		return nil
	}
	return db
}

func UserExist(db *sql.DB, userId uint64, apiKey string) bool {
	var exist bool
	err := db.QueryRow("select exists (select 1 from users where id = $1 and apikey = $2)", userId, apiKey).Scan(&exist)
	if err != nil {
		return false
	}
	return exist
}

func WriteUnpipedData(db *sql.DB, data UnpipedData) error {
	queryStr := "insert into unpiped_data (user_id, application, data) values ($1, $2, $3)"
	_, err := db.Exec(queryStr, data.UserId, data.Application, data.Data)
	return err
}

func SelectApplicationId(db *sql.DB, name string, companyId uint64) (applicationId uint64) {
	file, err := os.Open("database/select_application_id.sql")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()
	var fileBytes [256]byte
	_, err = file.Read(fileBytes[:])
	if err != nil {
		fmt.Println("Error reading file ", err)
		return
	}
	// queryStr := string(fileBytes[:])
	// fmt.Println("=== Here is the original queryyyyy: ", queryStr)
	// queryStr2 := strings.ReplaceAll(queryStr, "\n", " ")
	// fmt.Printf("=== Query after replace: ")
	// fmt.Printf("%s\n", queryStr2)
	// queryStr3 := strings.ReplaceAll(queryStr2, "\t", " ")
	// fmt.Println("=== Query after replace: ", queryStr3)

	queryStr := string(fileBytes[:])
	queryStr = strings.ReplaceAll(queryStr, "\n", " ")
	queryStr = strings.ReplaceAll(queryStr, "\t", " ")
	queryStr = strings.Join(strings.Fields(queryStr), " ")
	fmt.Println("=== Query after removing extra spaces: ", queryStr)

	fmt.Println("=== Query: ", queryStr, " name: ", name, " companyId: ", companyId)
	err = db.QueryRow(queryStr, name, companyId).Scan(&applicationId)
	// queryStr = "select id from applications where name = 'River level monitoring' and company_id = 1"
	// err = db.QueryRow(queryStr).Scan(&applicationId)
	if err != nil {
		fmt.Println("Error reading applicationId: ", err)
		applicationId = 0
	}
	return applicationId
}

func SelectApplicationDataStructure(db *sql.DB, applicationId uint64, version int) {
	file, err := os.Open("database/select_application_data_structures.sql")
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer file.Close()
	var fileBytes [256]byte
	_, err = file.Read(fileBytes[:])
	if err != nil {
		fmt.Println("Error reading file ", err)
		return
	}
	queryStr := string(fileBytes[:])
	fmt.Println("Query ", queryStr)
}
