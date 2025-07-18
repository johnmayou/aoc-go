package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	StateOn uint8 = iota
	StateOff
	StateOnToOff
	StateOffToOn
)

var NEIGHBOR_DIRECTIONS = [8][2]int8{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}} // lower left moving clockwise

func GridStep(grid [][]uint8) error {
	rows, cols := len(grid), len(grid[0])
	for r := range rows {
		for c := range cols {
			on := 0
			for i := range NEIGHBOR_DIRECTIONS {
				dir := NEIGHBOR_DIRECTIONS[i]
				dr, dc := dir[0], dir[1]
				nr, nc := r+int(dr), c+int(dc)
				if nr >= 0 && nr < rows &&
					nc >= 0 && nc < cols &&
					(grid[nr][nc] == StateOn ||
						grid[nr][nc] == StateOnToOff) {
					on += 1
				}
			}
			switch grid[r][c] {
			case StateOn:
				if !(on == 2 || on == 3) {
					grid[r][c] = StateOnToOff
				}
			case StateOff:
				if on == 3 {
					grid[r][c] = StateOffToOn
				}
			default:
				return fmt.Errorf("unexpected grid state: %d", grid[r][c])
			}
		}
	}
	for r := range grid {
		for c := range grid[r] {
			switch grid[r][c] {
			case StateOn, StateOff:
				// fallthrough
			case StateOnToOff:
				grid[r][c] = StateOff
			case StateOffToOn:
				grid[r][c] = StateOn
			default:
				return fmt.Errorf("unexpected grid state: %d", grid[r][c])
			}
		}
	}
	return nil
}

func GridCountOn(grid [][]uint8) int {
	count := 0
	for r := range len(grid) {
		for c := range len(grid[r]) {
			if grid[r][c] == StateOn {
				count += 1
			}
		}
	}
	return count
}

func ParseGrid(in io.Reader, rows int, cols int) ([][]uint8, error) {
	grid := make([][]uint8, rows)
	reader := bufio.NewReader(in)
	buf := make([]byte, 1)
	for r := range rows {
		row := make([]uint8, cols)
		for c := range cols {
			for {
				n, err := reader.Read(buf)
				if err != nil {
					if err == io.EOF {
						return nil, fmt.Errorf("reached EOF before reading %d x %d bytes", rows, cols)
					}
					return nil, fmt.Errorf("error reading byte: %s", err)
				}
				if n != 1 {
					return nil, fmt.Errorf("expected to read one byte but read (%d): %v", n, buf)
				}
				if buf[0] == '\n' {
					continue
				}
				break
			}
			switch buf[0] {
			case '#':
				row[c] = StateOn
			case '.':
				row[c] = StateOff
			default:
				return nil, fmt.Errorf("unexpected byte found: %v", buf[0])
			}
		}
		grid[r] = row
	}
	_, err := reader.Read(buf)
	if err != io.EOF {
		return nil, fmt.Errorf("expected to reach EOF after reading %d x %d bytes, but read: %v", rows, cols, buf)
	}
	return grid, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	grid, err := ParseGrid(file, 100, 100)
	if err != nil {
		log.Fatalf("error parsing grid: %s", err)
	}
	for range 100 {
		if err := GridStep(grid); err != nil {
			log.Fatalf("error with grid step: %s", err)
		}
	}
	fmt.Println(GridCountOn(grid))
}
