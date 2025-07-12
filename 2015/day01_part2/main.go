package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func FindFirstBasementPos(stream io.Reader) (int, error) {
	floor, pos := 0, 0
	reader := bufio.NewReader(stream)
	for {
		pos += 1
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

		if floor == -1 {
			return pos, nil
		}
	}
	return 0, fmt.Errorf("did not ever reach the basement")
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	pos, err := FindFirstBasementPos(file)
	if err != nil {
		log.Fatalf("failed to find first basement pos: %s", err)
	}
	fmt.Println(pos)
}
