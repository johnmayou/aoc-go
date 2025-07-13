package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func IsNice(str string) bool {
	vowelCnt := 0
	chTwiceInARow := false
	for i, ch := range str {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCnt += 1
		}
		if i+1 < len(str) {
			if str[i] == str[i+1] {
				chTwiceInARow = true
			}
			if (str[i] == 'a' && str[i+1] == 'b') ||
				(str[i] == 'c' && str[i+1] == 'd') ||
				(str[i] == 'p' && str[i+1] == 'q') ||
				(str[i] == 'x' && str[i+1] == 'y') {
				return false
			}
		}
	}
	return vowelCnt >= 3 && chTwiceInARow
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
