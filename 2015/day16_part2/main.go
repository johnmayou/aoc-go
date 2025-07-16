package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Sue struct {
	Number     uint
	Attributes map[string]uint8
}

type MFCSAM struct {
	Children    uint8
	Cats        uint8
	Samoyeds    uint8
	Pomeranians uint8
	Akitas      uint8
	Vizslas     uint8
	Goldfish    uint8
	Trees       uint8
	Cars        uint8
	Perfumes    uint8
}

func FindSue(sues []Sue, mfcsam MFCSAM) (int, error) {
	scores := make([]float64, len(sues))
	for i := range sues {
		score := 0.0
		for key, val := range sues[i].Attributes {
			switch key {
			case "children":
				score += calcScore(uintAbs(val, mfcsam.Children))
			case "cats":
				if val > mfcsam.Cats {
					score += 1.0
				}
			case "samoyeds":
				score += calcScore(uintAbs(val, mfcsam.Samoyeds))
			case "pomeranians":
				if val < mfcsam.Pomeranians {
					score += 1.0
				}
			case "akitas":
				score += calcScore(uintAbs(val, mfcsam.Akitas))
			case "vizslas":
				score += calcScore(uintAbs(val, mfcsam.Vizslas))
			case "goldfish":
				if val < mfcsam.Goldfish {
					score += 1.0
				}
			case "trees":
				if val > mfcsam.Trees {
					score += 1.0
				}
			case "cars":
				score += calcScore(uintAbs(val, mfcsam.Cars))
			case "perfumes":
				score += calcScore(uintAbs(val, mfcsam.Perfumes))
			default:
				return 0, fmt.Errorf("unexpected sue attribute: %q", key)
			}
			scores[i] = score
		}
	}
	if i := findMaxIndex(scores); i != -1 {
		return int(sues[i].Number), nil
	} else {
		return i, nil
	}
}

func uintAbs[T constraints.Unsigned](a, b T) T {
	if a > b {
		return a - b
	}
	return b - a
}

func calcScore[T constraints.Integer | constraints.Float](diff T) float64 {
	d := float64(diff)
	if d == 0 {
		return 1
	} else if d <= 1 {
		return 0.8
	} else if d <= 2 {
		return 0.6
	} else {
		return 0
	}
}

func findMaxIndex[T constraints.Ordered](arr []T) int {
	if len(arr) == 0 {
		return -1
	}
	maxVal := arr[0]
	maxIdx := 0
	for i := range arr {
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIdx = i
		}
	}
	return maxIdx
}

var sueRegex = regexp.MustCompile(`^Sue (\d+): (.+)$`)

func ParseSues(in io.Reader) ([]Sue, error) {
	var sues []Sue
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		matches := sueRegex.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, fmt.Errorf("invalid line format: %q", line)
		}
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing sue number: %s", err)
		}
		sue := Sue{
			Number:     uint(num),
			Attributes: make(map[string]uint8),
		}
		parts := strings.Split(matches[2], ", ")
		for _, part := range parts {
			kv := strings.Split(part, ": ")
			if len(kv) != 2 {
				return nil, fmt.Errorf("invalid attribute: %q", part)
			}
			val, err := strconv.ParseUint(kv[1], 10, 8)
			if err != nil {
				return nil, fmt.Errorf("error parsing value as uint8: %s", err)
			}
			sue.Attributes[kv[0]] = uint8(val)
		}
		sues = append(sues, sue)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %s", err)
	}
	return sues, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	sues, err := ParseSues(file)
	if err != nil {
		log.Fatalf("error parsing sues: %s", err)
	}
	mfcsam := MFCSAM{
		Children:    3,
		Cats:        7,
		Samoyeds:    2,
		Pomeranians: 3,
		Akitas:      0,
		Vizslas:     0,
		Goldfish:    5,
		Trees:       3,
		Cars:        2,
		Perfumes:    1,
	}
	sue, err := FindSue(sues, mfcsam)
	if err != nil {
		log.Fatalf("error finding sue: %s", err)
	}
	fmt.Println(sue)
}
