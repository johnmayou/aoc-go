package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func IsNice(str string) bool {
	pairs := make(map[string]int, len(str))
	pairRepeats := false
	pairWithOneInBetween := false
	for i := range str {
		if i+2 < len(str) && str[i] == str[i+2] {
			pairWithOneInBetween = true
		}
		if i+1 < len(str) {
			pair := str[i : i+2]
			matchingPairStart, exists := pairs[pair]
			if exists && matchingPairStart < i-1 { // make sure it didn't just start
				pairRepeats = true
			} else {
				pairs[pair] = i
			}
		}
	}
	return pairRepeats && pairWithOneInBetween
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	niceCnt := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatalf("error scanning text: %s", err)
		}
		if IsNice(line) {
			niceCnt += 1
		}
	}
	fmt.Println(niceCnt)
}
