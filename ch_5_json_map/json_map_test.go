package ch_5_json_map

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMap(t *testing.T) {
	jsonString := `{"id":"001","name":"Nurdin","age":20,"address":"JL. Pangeran Samudera","balance_count":9999999999999}`
	jsonByte := []byte(jsonString)

	var decoded map[string]any

	json.Unmarshal(jsonByte, &decoded)

	fmt.Println(decoded)
	fmt.Println(decoded["id"])
	fmt.Println(decoded["name"])
	fmt.Println(decoded["age"])
	fmt.Println(decoded["address"])
	fmt.Println(decoded["balance_count"])
}
