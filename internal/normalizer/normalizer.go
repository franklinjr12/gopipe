package normalizer

import (
	"database/sql"
	"encoding/json"
)

type ApplicationBytesDecode struct {
	FirstByte int
	LastByte  int
	Type      string
}

func ToJson(data []byte) (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	return jsonData, err
}

func RowsToFormatStruct(rows *sql.Rows) []ApplicationBytesDecode {
	if rows == nil {
		return nil
	}
	defer rows.Close()
	var formatArray []ApplicationBytesDecode
	for rows.Next() {
		var format ApplicationBytesDecode
		rows.Scan(&format.FirstByte, &format.LastByte, &format.Type)
		formatArray = append(formatArray, format)
	}
	return formatArray
}
