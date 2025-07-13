package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Coord struct {
	X int
	Y int
}

func CountUniqueDeliveries(stream io.Reader) (int, error) {
	houses := make(map[Coord]struct{})
	x, y := 1, 1
	houses[Coord{X: x, Y: y}] = struct{}{}

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
		case '>':
			x += 1
		case '<':
			x -= 1
		case '^':
			y += 1
		case 'v':
			y -= 1
		default:
			return 0, fmt.Errorf("unexpected byte: %s", err)
		}
		houses[Coord{X: x, Y: y}] = struct{}{}
	}
	return len(houses), nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	houses, err := CountUniqueDeliveries(file)
	if err != nil {
		log.Fatalf("error counting deliveries: %s", err)
	}
	fmt.Println(houses)
}
