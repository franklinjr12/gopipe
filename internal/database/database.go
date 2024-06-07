package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type UnpipedData struct {
	Data        []byte
	Application string
	UserId      uint64
}

type PipedData struct {
	Id            uint64
	ApplicationId uint64
	Version       int
	Data          []byte
	CreatedAt     string
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
	queryFile := "select_application_id.sql"
	queryStr, err := readQueryFile(queryFile)
	if err != nil {
		fmt.Println("Error reading query file ", queryFile, " ", err)
		return 0
	}
	err = db.QueryRow(queryStr, name, companyId).Scan(&applicationId)
	if err != nil {
		fmt.Println("Error reading applicationId: ", err)
		return 0
	}
	return applicationId
}

func SelectApplicationDataStructure(db *sql.DB, applicationId uint64, version int) *sql.Rows {
	queryFile := "select_application_data_structures.sql"
	queryStr, err := readQueryFile(queryFile)
	if err != nil {
		fmt.Println("Error reading query file ", queryFile, " ", err)
		return nil
	}
	rows, err := db.Query(queryStr, applicationId, version)
	if err != nil {
		fmt.Println("Error reading application data structure ", err)
		return nil
	}
	return rows
}

func readQueryFile(fileName string) (string, error) {
	file, err := os.Open("database/" + fileName)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return "", err
	}
	defer file.Close()
	var fileBytes [256]byte
	var bytesRead int
	bytesRead, err = file.Read(fileBytes[:])
	if err != nil {
		fmt.Println("Error reading file ", err)
		return "", err
	}
	queryStr := string(fileBytes[:bytesRead])
	return queryStr, nil
}

func InsertPipedData(db *sql.DB, pipedData *PipedData) error {
	queryFile := "insert_piped_data.sql"
	queryStr, err := readQueryFile(queryFile)
	if err != nil {
		fmt.Println("Error reading query file ", queryFile, " ", err)
		return nil
	}
	_, err = db.Exec(queryStr, pipedData.ApplicationId, pipedData.Version, pipedData.Data)
	return err
}
