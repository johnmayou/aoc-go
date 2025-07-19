package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

type Replacement struct {
	From string
	To   string
}

func FastestCreationSteps(replacements []Replacement, target string) int {
	revReplacements := make([]Replacement, len(replacements))
	for i := range replacements {
		revReplacements[i] = Replacement{
			From: replacements[i].To,
			To:   replacements[i].From,
		}
	}
	slices.SortFunc(revReplacements, func(a, b Replacement) int {
		if len(a.From) > len(b.From) {
			return -1
		} else if len(a.From) < len(b.From) {
			return 1
		} else {
			return 0
		}
	})
	steps := 0
	molecule := target
	for molecule != "e" {
		for i := range revReplacements {
			from, to := revReplacements[i].From, revReplacements[i].To
			if strings.Contains(molecule, from) {
				molecule = strings.Replace(molecule, from, to, 1)
				steps += 1
				break
			}
		}
	}
	return steps
}

func Parse(in io.Reader) (replacements []Replacement, molecule string, err error) {
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
		replacements = append(replacements, Replacement{
			From: parts[0],
			To:   parts[1],
		})
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
	fmt.Println(FastestCreationSteps(replacements, molecule))
}
