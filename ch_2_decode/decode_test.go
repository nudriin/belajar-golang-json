package ch_2_decode

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	Name    string
	Age     int
	Address string
	Married bool
	Hobbies []string
}

func Encode(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return bytes
}

func Decode(data []byte) *Person {
	// ! Decode pada golang dapat menggunakan  function Unmarshal
	// ! returnnya adalah  error
	// ! parameternya ada (byte[], interface{})
	// * Byte adalah data json yang akan di decode
	// * Interface adalah tempat menyimpan data, berupa pointer untuk passing by reference

	person := &Person{}

	err := json.Unmarshal(data, person)
	if err != nil {
		panic(err)
	}

	return person
}

func TestDecode(t *testing.T) {
	jsonData := Encode(Person{
		Name:    "Nurdin",
		Age:     20,
		Address: "JL. Pangeran Samudera",
		Married: false,
		Hobbies: []string{"Ngoding", "Game Building", "Web Building"},
	})

	decoded := Decode(jsonData)

	fmt.Println(decoded.Name)
	fmt.Println(decoded.Hobbies)
	fmt.Println(decoded.Age)
}
