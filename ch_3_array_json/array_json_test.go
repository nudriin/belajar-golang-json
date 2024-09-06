package ch_3_array_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street  string
	City    string
	Provice string
	Country string
}

type Person struct {
	Name      string
	Age       int
	Hobbies   []string
	Addresses []Address // Membuat struct dengan data yang komplex
}

func Encode(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return bytes
}

func Decode(data []byte) *Person {

	person := &Person{}

	err := json.Unmarshal(data, person)
	if err != nil {
		panic(err)
	}

	return person
}

func TestArrayJson(t *testing.T) {
	jsonData := Encode(Person{
		Name: "Nurdin",
		Age:  20,
		Hobbies: []string{
			"Coding",
			"Making Game",
			"Making Web",
		},
		Addresses: []Address{
			{
				Street:  "JL. Pangeran Samudera",
				City:    "Palangka Raya",
				Provice: "Central Kalimantan",
				Country: "Indonesia",
			},
			{
				Street:  "JL. Setia Yakin",
				City:    "Sukamara",
				Provice: "Central Kalimantan",
				Country: "Indonesia",
			},
		},
	})

	decoded := Decode(jsonData)

	fmt.Println(decoded)
	fmt.Println(decoded.Hobbies[0])
	fmt.Println(decoded.Addresses)
	fmt.Println(decoded.Addresses[0])
	fmt.Println(decoded.Addresses[1])
}
