package ch_6_stream_encoder

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	writer, _ := os.Create("data.json")
	decoder := json.NewEncoder(writer)

	person := map[string]any{
		"name": "Nurdin",
		"age":  20,
	}

	// ! Cukup encode menggunakan function Encode
	decoder.Encode(person)
}
