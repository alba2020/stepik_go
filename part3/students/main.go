package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName string
	FirstName string
	MiddleName string
	Birthday string
	Address string
	Phone string
	Rating []int
}

type Group struct {
	ID int
	Number string
	Year int
	Students []Student
}

func main() {
	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	var group Group

	if err := json.Unmarshal(data, &group); err != nil {
		fmt.Println(err)
		return
	}

	total := 0
	for _, student := range group.Students {
		total += len(student.Rating)
	}

	res := make(map[string]float64)
	res["Average"] = float64(total) / float64(len(group.Students))

	dat, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", dat)
}
