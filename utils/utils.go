package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadAyatData() map[string][]string {
	// Load ayat from the ayat.json file
	file, err := ioutil.ReadFile("data/ayat.json")
	if err != nil {
		log.Fatalf("Error reading ayat.json: %v", err)
	}

	var ayatData map[string][]string
	err = json.Unmarshal(file, &ayatData)
	if err != nil {
		log.Fatalf("Error unmarshalling ayat.json: %v", err)
	}

	return ayatData
}
