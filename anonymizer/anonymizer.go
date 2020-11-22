package anonymizer

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const (
	replacementDataDir = "../replacement_data/"
	addressesFile      = "addresses.txt"
	firstNamesFile     = "first_names.txt"
	lastNamesFile      = "last_names.txt"
)

var (
	replacementAddresses  []string
	replacementFirstNames []string
	replacementLastNames  []string
)

func init() {
	loadReplacementData()

	rand.Seed(time.Now().UnixNano())
}

func loadReplacementData() error {
	addresses, err := ioutil.ReadFile(replacementDataDir + addressesFile)
	if err != nil {
		return err
	}

	replacementAddresses = strings.Split(string(addresses), "\n")

	firstNames, err := ioutil.ReadFile(replacementDataDir + firstNamesFile)
	if err != nil {
		return err
	}

	replacementFirstNames = strings.Split(string(firstNames), "\n")

	lastNames, err := ioutil.ReadFile(replacementDataDir + lastNamesFile)
	if err != nil {
		return err
	}

	replacementLastNames = strings.Split(string(lastNames), "\n")

	return nil
}

func GetAnonymousValues() map[string]json.RawMessage {
	anonData := make(map[string]json.RawMessage, 5)

	r := rand.Intn(len(replacementFirstNames))
	anonData["first_name"] = []byte("\"" + replacementFirstNames[r] + "\"")

	return anonData
}
