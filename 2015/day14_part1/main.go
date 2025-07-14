package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Reindeer struct {
	Name        string
	FlySpeed    uint8
	FlySeconds  uint8
	RestSeconds uint8
}

func Race(reindeer []Reindeer, seconds uint16) (string, int) {
	distances := make([]int, 0, len(reindeer))
	for i := range reindeer {
		distances = append(distances, calcDistance(reindeer[i], seconds))
	}
	winnerIdx := findMaxIndex(distances)
	return reindeer[winnerIdx].Name, distances[winnerIdx]
}

func calcDistance(reindeer Reindeer, seconds uint16) int {
	dist := 0
	flyLeft := reindeer.FlySeconds
	for seconds > 0 {
		if flyLeft > 0 {
			dist += int(reindeer.FlySpeed)
			flyLeft -= 1
			seconds -= 1
		} else {
			if seconds < uint16(reindeer.RestSeconds) {
				break
			}
			seconds -= uint16(reindeer.RestSeconds)
			flyLeft = reindeer.FlySeconds
		}
	}
	return dist
}

func findMaxIndex(arr []int) int {
	switch len(arr) {
	case 0:
		return -1
	case 1:
		return 0
	default:
		maxVal := arr[0]
		maxIdx := 0
		for i := 1; i < len(arr); i++ {
			if arr[i] > maxVal {
				maxVal = arr[i]
				maxIdx = i
			}
		}
		return maxIdx
	}
}

func ParseReindeer(in io.Reader) ([]Reindeer, error) {
	var reindeer []Reindeer
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error scanning line: %s", err)
		}
		matches := regexp.MustCompile(`^(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.$`).FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("invalid reindeer format: %s", line)
		}
		flySpeed, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing fly speed (%s): %s", line, err)
		}
		flySeconds, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing fly seconds (%s): %s", line, err)
		}
		restSeconds, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, fmt.Errorf("error parsing rest seconds (%s): %s", line, err)
		}
		reindeer = append(reindeer, Reindeer{
			Name:        matches[1],
			FlySpeed:    uint8(flySpeed),
			FlySeconds:  uint8(flySeconds),
			RestSeconds: uint8(restSeconds),
		})
	}
	return reindeer, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	reindeer, err := ParseReindeer(file)
	if err != nil {
		log.Fatalf("error parsing reindeer: %s", err)
	}
	winner, distance := Race(reindeer, 2503)
	fmt.Printf("%s won with a distance of %d\n", winner, distance)
}
