package anonymizer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	replacementDataDir = "./replacement_data/"
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
	var err error

	replacementAddresses, err = stringsFromFile(addressesFile)
	if err != nil {
		return err
	}

	replacementFirstNames, err = stringsFromFile(firstNamesFile)
	if err != nil {
		return err
	}

	replacementLastNames, err = stringsFromFile(lastNamesFile)

	return err
}

// GetAnonymousValues returns a map of sensitive field names to randomized values
func GetAnonymousValues() map[string]json.RawMessage {
	anonData := make(map[string]json.RawMessage, 5)

	r := rand.Intn(len(replacementFirstNames))
	anonData["first_name"] = []byte("\"" + replacementFirstNames[r] + "\"")

	r = rand.Intn(len(replacementLastNames))
	anonData["last_name"] = []byte("\"" + replacementLastNames[r] + "\"")

	r = rand.Intn(len(replacementAddresses))
	anonData["address"] = []byte("\"" + replacementAddresses[r] + "\"")

	anonData["dob"] = []byte("\"" + randDOB() + "\"")

	anonData["phone"] = []byte("\"" + randPhone() + "\"")

	return anonData
}

// returns a random birth date in the range of (today - 100 years) to (today - 18 years)
func randDOB() string {
	min := time.Now().AddDate(-100, 0, 0).Unix()
	max := time.Now().AddDate(-18, 0, 0).Unix()
	birthTime := time.Unix(rand.Int63n(max-min)+min, 0)

	return birthTime.Format("2006-01-02")
}

func randPhone() string {
	// generate a number from 0 to 9
	firstDigit := strconv.Itoa(rand.Intn(8) + 1)

	// generate a 9-digit number and pad with leading 0s if necessary
	last9Digits := fmt.Sprintf("%09d", rand.Intn(1000000000))

	return firstDigit + last9Digits
}

func stringsFromFile(filename string) ([]string, error) {
	bytes, err := ioutil.ReadFile(replacementDataDir + filename)
	if err != nil {
		return nil, err
	}

	// remove \r characters
	s := strings.ReplaceAll(string(bytes), "\r", "")

	// split each line into a separate string
	return strings.Split(s, "\n"), nil
}
