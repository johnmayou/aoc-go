package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Reindeer struct {
	Name        string
	FlySpeed    uint8
	FlySeconds  uint8
	RestSeconds uint8
}

func Race(reindeer []Reindeer, seconds uint16) (int, error) {
	raceState := make([]*Reindeer, len(reindeer))
	for i := range reindeer {
		cpy := reindeer[i]
		raceState[i] = &cpy
	}
	distances, points := make([]int, len(reindeer)), make([]int, len(reindeer))
	for range seconds {
		for i := range reindeer {
			rein := raceState[i]
			if rein.FlySeconds > 0 {
				distances[i] += int(rein.FlySpeed)
				rein.FlySeconds--
				if rein.FlySeconds == 0 {
					rein.RestSeconds = reindeer[i].RestSeconds
				}
			} else if rein.RestSeconds > 0 {
				rein.RestSeconds--
				if rein.RestSeconds == 0 {
					rein.FlySeconds = reindeer[i].FlySeconds
				}
			} else {
				return 0, fmt.Errorf("expected reindeer #%d to have at fly seconds or rest seconds left at %d seconds left: %+v", i, seconds, *rein)
			}
		}
		maxVal := distances[0]
		maxIdxs := []int{0}
		for i := 1; i < len(distances); i++ {
			if distances[i] > maxVal {
				maxVal = distances[i]
				maxIdxs = maxIdxs[:0]
				maxIdxs = append(maxIdxs, i)
			} else if distances[i] == maxVal {
				maxIdxs = append(maxIdxs, i)
			}
		}
		for _, i := range maxIdxs {
			points[i]++
		}
	}
	return slices.Max(points), nil
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
	distance, err := Race(reindeer, 2503)
	if err != nil {
		log.Fatalf("error racing: %s", err)
	}
	println(distance)
}
