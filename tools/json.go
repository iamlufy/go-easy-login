package tools

import (
	"encoding/json"
	"fmt"
)

func JsonString(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		panic(fmt.Errorf(err.Error(), value))
	}
	return string(json)
}

func JsonStringToMap(jsonString string) map[string]interface{} {
	var f interface{}
	err := json.Unmarshal([]byte(jsonString), &f)
	if err != nil {
		panic(fmt.Errorf(err.Error(), jsonString))
	}
	return f.(map[string]interface{})
}
