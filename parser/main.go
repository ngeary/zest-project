package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/ngeary/zest-project/anonymizer"
)

const (
	dataDir     = "../data/"
	anonDataDir = "../anon_data/"
)

// Request represents a json request to insert data
type Request struct {
	RequestID string `json:"request_id"`
	Rows      []*Row
}

// Row struct
type Row struct {
	RowID   string    `json:"row_id"`
	Sources []*Source `json:"sources"`
}

// Source struct
type Source struct {
	Name    string      `json:"name"`
	Version int         `json:"version"`
	Format  string      `json:"format"`
	Values  interface{} `json:"values"`
}

func main() {
	seenFiles := make(map[string]bool)

	for {
		fileInfos, err := ioutil.ReadDir(dataDir)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, fi := range fileInfos {
			if seenFiles[fi.Name()] || !strings.HasSuffix(strings.ToLower(fi.Name()), ".json") {
				continue
			}

			err = parse(fi.Name())
			if err != nil {
				log.Println(err)
			}

			seenFiles[fi.Name()] = true
		}

		time.Sleep(time.Second)
	}
}

func parse(filename string) error {
	file, err := ioutil.ReadFile(dataDir + filename)
	if err != nil {
		return err
	}

	file = removeXMLDeclarations(file)

	req := Request{}

	err = json.Unmarshal(file, &req)
	if err != nil {
		return err
	}

	// anonymize some of the applicant data
	for _, row := range req.Rows {
		for _, source := range row.Sources {
			if source.Name != "app_data" {
				continue
			}

			var vals map[string]json.RawMessage

			switch strings.ToLower(source.Format) {
			case "json":
				vals, err = jsonToMap(source.Values)
			case "csv":
				vals, err = csvToMap(source.Values)
			case "xml":
				vals, err = xmlToMap(source.Values)
			default:
				err = errors.New("unrecognized data format: " + source.Format)
			}

			if err != nil {
				return err
			}

			anonData := anonymizer.GetAnonymousValues()
			for k, v := range anonData {
				vals[k] = v
			}

			source.Values = vals
		}
	}

	return writeToFile(&req, filename)
}

func writeToFile(r *Request, filename string) error {
	bytes, err := json.MarshalIndent(r, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(anonDataDir+fmt.Sprintf("%d-", time.Now().UnixNano())+filename, bytes, 0644)
}
