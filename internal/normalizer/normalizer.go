package normalizer

import "encoding/json"

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
