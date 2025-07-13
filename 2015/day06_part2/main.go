package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Action int

const (
	TurnOn Action = iota
	TurnOff
	Toggle
)

type Coord struct {
	X int
	Y int
}

type Instruction struct {
	Action
	Start Coord
	Stop  Coord
}

func Perform(grid [][]int, instructions []Instruction) {
	for _, instr := range instructions {
		for x := instr.Start.X; x <= instr.Stop.X; x++ {
			for y := instr.Start.Y; y <= instr.Stop.Y; y++ {
				switch instr.Action {
				case TurnOn:
					grid[x][y] += 1
				case TurnOff:
					grid[x][y] = max(0, grid[x][y]-1)
				case Toggle:
					grid[x][y] += 2
				}
			}
		}
	}
}

func CountBrightness(grid [][]int) int {
	count := 0
	for row := range grid {
		for col := range grid[row] {
			count += grid[row][col]
		}
	}
	return count
}

func ParseInstructions(stream io.Reader) ([]Instruction, error) {
	var instructions []Instruction
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error scanning line: %s", err)
		}
		var action Action
		if strings.HasPrefix(line, "turn on") {
			action = TurnOn
			line = line[8:]
		} else if strings.HasPrefix(line, "turn off") {
			action = TurnOff
			line = line[9:]
		} else if strings.HasPrefix(line, "toggle") {
			action = Toggle
			line = line[7:]
		} else {
			return nil, fmt.Errorf("invalid line action: %s", line)
		}
		parts := strings.Split(line, " through ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid coordinates: %s", line)
		}
		startCoord, err := parseStrCoord(parts[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing coordinate: %s", err)
		}
		stopCoord, err := parseStrCoord(parts[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing coordinate: %s", err)
		}
		instructions = append(instructions, Instruction{
			Action: action,
			Start:  startCoord,
			Stop:   stopCoord,
		})
	}
	return instructions, nil
}

func parseStrCoord(str string) (Coord, error) {
	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		return Coord{}, fmt.Errorf("invalid format: %s", str)
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Coord{}, fmt.Errorf("invalid format: %s", str)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Coord{}, fmt.Errorf("invalid format: %s", str)
	}
	return Coord{X: x, Y: y}, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	instructions, err := ParseInstructions(file)
	if err != nil {
		log.Fatalf("error parsing instructions: %s", err)
	}
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	Perform(grid, instructions)
	fmt.Println(CountBrightness(grid))
}
