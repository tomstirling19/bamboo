package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
)

func ToJSONString(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("Failed to convert to JSON string: %v", err)
		return ""
	}

	var compacted bytes.Buffer
	if err := json.Compact(&compacted, data); err != nil {
		log.Printf("Failed to compact JSON string: %v", err)
		return string(data) 
	}

	prettyStr := cleanJSON(compacted.String())

	return prettyStr
}

func cleanJSON(input string) string {
	re := regexp.MustCompile(`(?s:\\n|\\t|\s{2,}|\t|\n)+`)
	return re.ReplaceAllString(input, " ")
}
