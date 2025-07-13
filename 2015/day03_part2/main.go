package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type SantaType int

const (
	NormalSanta = iota
	RobotSanta
)

type Coord struct {
	X int
	Y int
}

func CountUniqueDeliveries(stream io.Reader) (int, error) {
	houses := make(map[Coord]struct{})
	sx, sy, rx, ry := 1, 1, 1, 1
	houses[Coord{X: 1, Y: 1}] = struct{}{}
	santaTurn := NormalSanta

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
			if santaTurn == NormalSanta {
				sx += 1
			} else {
				rx += 1
			}
		case '<':
			if santaTurn == NormalSanta {
				sx -= 1
			} else {
				rx -= 1
			}
		case '^':
			if santaTurn == NormalSanta {
				sy += 1
			} else {
				ry += 1
			}
		case 'v':
			if santaTurn == NormalSanta {
				sy -= 1
			} else {
				ry -= 1
			}
		default:
			return 0, fmt.Errorf("unexpected byte: %s", err)
		}
		if santaTurn == NormalSanta {
			houses[Coord{X: sx, Y: sy}] = struct{}{}
			santaTurn = RobotSanta
		} else {
			houses[Coord{X: rx, Y: ry}] = struct{}{}
			santaTurn = NormalSanta
		}
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
