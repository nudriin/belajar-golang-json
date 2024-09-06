package ch_6_stream_decoder

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("./data.json")
	decoder := json.NewDecoder(reader)

	var result map[string]any

	// ! Decode datanya cukup gunakan function Decode
	decoder.Decode(&result)

	fmt.Println(result)
	fmt.Println(result["name"])
	fmt.Println(result["age"])
}
