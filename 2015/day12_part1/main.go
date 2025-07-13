package main

import (
	"encoding/json"
	"log"
	"os"
)

func AddAllNumbers(data any) int {
	switch v := data.(type) {
	case float64:
		return int(v)
	case []any:
		sum := 0
		for _, el := range v {
			sum += AddAllNumbers(el)
		}
		return sum
	case map[string]any:
		sum := 0
		for _, val := range v {
			sum += AddAllNumbers(val)
		}
		return sum
	default:
		return 0
	}
}

func main() {
	rawJson, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	var jsonData any
	if err = json.Unmarshal(rawJson, &jsonData); err != nil {
		log.Fatalf("error unmarshaling json data: %s", err)
	}
	println(AddAllNumbers(jsonData))
}
