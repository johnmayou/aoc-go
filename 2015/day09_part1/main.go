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

type Distance struct {
	From string
	To   string
	Len  int
}

type AdjDistance struct {
	To  string
	Len int
}

func FindShortestRoute(distances []Distance) int {
	adj := adjFromDistances(distances)
	places := make([]string, 0, len(adj))
	for place := range adj {
		places = append(places, place)
	}
	shortest := int(math.Inf(1))
	var dfs func(curr string, total int, visited map[string]struct{})
	dfs = func(curr string, total int, visited map[string]struct{}) {
		if total > shortest {
			return
		}
		if len(visited) == len(places) {
			if total < shortest {
				shortest = total
			}
			return
		}
		for _, next := range adj[curr] {
			if _, ok := visited[next.To]; ok {
				continue
			}
			visited[next.To] = struct{}{}
			dfs(next.To, total+next.Len, visited)
			delete(visited, next.To)
		}
	}
	visited := make(map[string]struct{})
	for _, start := range places {
		visited[start] = struct{}{}
		dfs(start, 0, visited)
		delete(visited, start)
	}
	return shortest
}

func adjFromDistances(distances []Distance) map[string][]AdjDistance {
	adj := make(map[string][]AdjDistance)
	for i := range distances {
		from, to, l := distances[i].From, distances[i].To, distances[i].Len
		if _, ok := adj[from]; !ok {
			adj[to] = []AdjDistance{}
		}
		if _, ok := adj[to]; !ok {
			adj[to] = []AdjDistance{}
		}
		adj[from] = append(adj[from], AdjDistance{To: to, Len: l})
		adj[to] = append(adj[to], AdjDistance{To: from, Len: l})
	}
	return adj
}

func ParseDistances(stream io.Reader) ([]Distance, error) {
	var distances []Distance
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error scanning line: %s", err)
		}
		matches := regexp.MustCompile(`^(\w+) to (\w+) = (\d+)$`).FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("invalid distance format: %s", line)
		}
		dist, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("error parsing distance (%s): %s", line, err)
		}
		distances = append(distances, Distance{
			From: matches[1],
			To:   matches[2],
			Len:  dist,
		})
	}
	return distances, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	distances, err := ParseDistances(file)
	if err != nil {
		log.Fatalf("error parsing distances: %s", err)
	}
	println(FindShortestRoute(distances))
}
