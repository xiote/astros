package main

import (
	"encoding/json"
)

func Decode(jsonData string, target interface{}) error {
	return json.Unmarshal([]byte(jsonData), target)
}
