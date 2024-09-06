package ch_1_encode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func EncodingToJson(data any) {
	// ! Encode data ke jsin menggunakan function marshal
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes)) //! balikannya berupa bytes
}

type Person struct {
	Name    string
	Age     int
	Address string
	Married bool
}

func TestEncode(t *testing.T) {
	EncodingToJson("Nurdin")
	EncodingToJson(1)
	EncodingToJson(true)

	EncodingToJson([]string{"Nurdin", "Hishasy", "Sunny", "Summer"})
	EncodingToJson(map[string]any{
		"Name":    "Nurdin",
		"Age":     20,
		"Address": "JL. Setia Yakin",
		"Married": false,
	})

	// ! Sebaiknya gunakan struct untuk datanya
	EncodingToJson(Person{
		Name:    "Nurdin",
		Age:     20,
		Address: "JL. Pangeran Samudera",
		Married: false,
	})
}
