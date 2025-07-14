package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Happiness struct {
	Person    string
	Neighbor  string
	Happiness int8
}

func FindOptimalSeating(scores []Happiness) (int, error) {
	hMap := make(map[string]int8)
	people := make(map[string]struct{})
	for i := range scores {
		happ := scores[i]
		hMap[happ.Person+"|"+happ.Neighbor] = happ.Happiness
		people[happ.Person] = struct{}{}
	}
	maximum := int(math.Inf(-1))
	var dfs func(seating []string, visited map[string]struct{}) error
	dfs = func(seating []string, visited map[string]struct{}) error {
		if len(visited) == len(people) {
			happ := 0
			for i := range len(seating) {
				key1 := seating[i] + "|" + seating[(i+1)%len(seating)]
				key2 := seating[(i+1)%len(seating)] + "|" + seating[i]
				if h, ok := hMap[key1]; ok {
					happ += int(h)
				} else {
					return fmt.Errorf("expected to find key '%s' in happiness map", key1)
				}
				if h, ok := hMap[key2]; ok {
					happ += int(h)
				} else {
					return fmt.Errorf("expected to find key '%s' in happiness map", key2)
				}
			}
			if happ > maximum {
				maximum = happ
			}
			return nil
		}
		for person := range people {
			if _, ok := visited[person]; ok {
				continue
			}
			seating = append(seating, person)
			visited[person] = struct{}{}
			if err := dfs(seating, visited); err != nil {
				return err
			}
			seating = seating[:len(seating)-1]
			delete(visited, person)
		}
		return nil
	}
	if err := dfs(make([]string, 0, len(people)), make(map[string]struct{})); err != nil {
		return 0, fmt.Errorf("error while recursing: %s", err)
	}
	return maximum, nil
}

func ParseHappinessScores(in io.Reader) ([]Happiness, error) {
	var scores []Happiness
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error scanning line: %s", err)
		}
		matches := regexp.MustCompile(`^(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+).$`).FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("invalid happiness score format: %s", line)
		}
		i, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing happiness (%s): %s", line, err)
		}
		if i < -128 || i > 127 {
			return nil, fmt.Errorf("value of of int8 range: %d", i)
		}
		var happiness int8
		switch matches[2] {
		case "gain":
			happiness = int8(i)
		case "lose":
			happiness = -int8(i)
		default:
			return nil, fmt.Errorf("invalid happiness direction: %s", line)
		}
		scores = append(scores, Happiness{
			Person:    matches[1],
			Neighbor:  matches[4],
			Happiness: happiness,
		})
	}
	return scores, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	scores, err := ParseHappinessScores(file)
	if err != nil {
		log.Fatalf("error parsing hapiness scores: %s", err)
	}
	happiness, err := FindOptimalSeating(scores)
	if err != nil {
		log.Fatalf("error finding optimal seating: %s", err)
	}
	println(happiness)
}
