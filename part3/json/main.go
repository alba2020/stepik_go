package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Item struct {
	GlobalId int64 `json:"global_id"`
	// SystemObjectId string `json:"-"`
	// Kod string `json:"-"`
	// IsDeleted int `json:"-"`
	// SignatureDate string `json:"-"`
	// Nomdescr string `json:"-"`
	// Idx string `json:"-"`
	// Razdel string `json:"-"`
	// Name string `json:"-"`
}

func main() {
	jsonFile, err := os.Open("./data-20190514T0100.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var items []Item

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'items' which we defined above
	json.Unmarshal(byteValue, &items)

	fmt.Println("total items", len(items))

	var sum int64 = 0
	for _, item := range items {
		sum += item.GlobalId
	}
	fmt.Println(sum)
}
