package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func FindFloor(stream io.Reader) (int, error) {
	floor := 0
	reader := bufio.NewReader(stream)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, fmt.Errorf("error reading byte: %s", err)
		}
		switch b {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		default:
			return 0, fmt.Errorf("unexpected byte read: %s", err)
		}
	}
	return floor, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	floor, err := FindFloor(file)
	if err != nil {
		log.Fatalf("failed to find floor: %s", err)
	}
	fmt.Println(floor)
}
