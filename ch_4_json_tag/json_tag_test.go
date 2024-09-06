package ch_3_array_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person struct {
	// ! Digunakan untuk mengubah penulisan nama dari attributenya
	Id           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Address      string `json:"address"`
	BalanceCount int    `json:"balance_count"`
}

type PersonNoTag struct {
	Id           string
	Name         string
	Age          int
	Address      string
	BalanceCount int
}

func Encode(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return bytes
}

func TestArrayJson(t *testing.T) {
	// ! Tanpa menggunakan TAG
	jsonData := string(Encode(PersonNoTag{
		Id:           "001",
		Name:         "Nurdin",
		Age:          20,
		Address:      "JL. Pangeran Samudera",
		BalanceCount: 9999999999999,
	}))
	fmt.Println(jsonData)

	// ! Menggunakan TAG
	jsonData = string(Encode(Person{
		Id:           "001",
		Name:         "Nurdin",
		Age:          20,
		Address:      "JL. Pangeran Samudera",
		BalanceCount: 9999999999999,
	}))
	fmt.Println(jsonData)
}

func TestDecodeTag(t *testing.T) {
	// * Jika datanya menggunakan format yang sama dengan tag maka akan terbaca datanya
	jsonData := `{"id":"001","name":"Nurdin","age":20,"address":"JL. Pangeran Samudera","balance_count":9999999999999}`
	byteJson := []byte(jsonData)

	decoded := &Person{}

	json.Unmarshal(byteJson, decoded)

	fmt.Println(decoded)

	// * Jika datanya menggunakan format yang TIDAK sama dengan tag maka akan terbaca datanya
	jsonData = `{"Id":"001","Name":"Nurdin","Age":20,"Address":"JL. Pangeran Samudera","BalanceCount":9999999999999}`
	byteJson = []byte(jsonData)

	decoded = &Person{}

	json.Unmarshal(byteJson, decoded)

	fmt.Println(decoded) //! BalanceCount tidak terbaca
}
