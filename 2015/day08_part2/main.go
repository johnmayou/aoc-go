package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func StringEncodingLengthDiff(rawStr string) int {
	return len(strconv.Quote(rawStr)) - len(rawStr)
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
		totalDiff += StringEncodingLengthDiff(line)
	}
	println(totalDiff)
}
