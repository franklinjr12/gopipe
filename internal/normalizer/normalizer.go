package normalizer

import (
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"math"
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

func BytesToStruct(data []byte, dataSize int, format []ApplicationBytesDecode) []any {
	results := make([]any, dataSize)
	elementSize := len(data) / dataSize
	for i := 0; i < dataSize; i++ {
		p := i * elementSize
		packet := data[p : p+elementSize]
		for _, v := range format {
			switch v.Type {
			case "int":
				results = append(results, int32(binary.LittleEndian.Uint32(packet[v.FirstByte:v.LastByte+1])))
			case "float":
				results = append(results, math.Float32frombits(binary.LittleEndian.Uint32(packet[v.FirstByte:v.LastByte+1])))
			case "bytes":
				results = append(results, string(packet[v.FirstByte:v.LastByte+1]))
			}
		}
	}
	return results

}
