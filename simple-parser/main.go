package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type testData struct {
	Nums    numFields    `json:"num_fields"`
	Strings stringFields `json:"string_fields"`
}

type numFields struct {
	N1 int `json:"n1"`
	N2 int `json:"n2"`
	N3 int `json:"n3"`
}

type stringFields struct {
	S1 string `json:"s1"`
	S2 string `json:"s2"`
}

func main() {
	data := testData{}

	file, err := ioutil.ReadFile("/Users/nick/github-repos/zest-project/data/ngtest1.json")
	if err != nil {
		log.Fatalf("error reading file: %v\n", err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("error unmarshaling data: %v\n", err)
	}

	log.Println(data)
}
