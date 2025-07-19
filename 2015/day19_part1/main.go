package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func DistinctMolecules(replacements map[string][]string, molecule string) int {
	molecules := mapset.NewSet[string]()
	for from, tos := range replacements {
		for i := 0; i <= len(molecule)-len(from); i++ {
			if molecule[i:i+len(from)] == from {
				for _, to := range tos {
					molecules.Add(molecule[:i] + to + molecule[i+len(from):])
				}
			}
		}
	}
	return molecules.Cardinality()
}

func Parse(in io.Reader) (replacements map[string][]string, molecule string, err error) {
	replacements = make(map[string][]string)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			return nil, "", fmt.Errorf("invalid replacement line: %s", line)
		}
		from, to := parts[0], parts[1]
		if _, ok := replacements[from]; !ok {
			replacements[from] = make([]string, 0)
		}
		replacements[from] = append(replacements[from], to)
	}
	if scanner.Scan() {
		molecule = scanner.Text()
	}
	if err = scanner.Err(); err != nil {
		return nil, "", err
	}
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	replacements, molecule, err := Parse(file)
	if err != nil {
		log.Fatalf("error parsing file: %s", err)
	}
	fmt.Println(DistinctMolecules(replacements, molecule))
}
