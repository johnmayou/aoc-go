package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func StringLiteralToInMemoryDiff(rawStr string) (int, error) {
	memStr, err := strconv.Unquote(rawStr)
	if err != nil {
		return 0, fmt.Errorf("error unquoting string: %s", err)
	}
	return len(rawStr) - len(memStr), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	totalDiff := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatalf("error scanning line: %s", err)
		}
		diff, err := StringLiteralToInMemoryDiff(line)
		if err != nil {
			log.Fatalf("error calculating diff: %s", err)
		}
		totalDiff += diff
	}
	println(totalDiff)
}
