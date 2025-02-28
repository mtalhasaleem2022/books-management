package utils

import "encoding/json"

func ToJSON(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

func FromJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
