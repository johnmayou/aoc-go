package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Package struct {
	L int
	W int
	H int
}

func TotalRibbon(packages []Package) int {
	total := 0
	for i := range packages {
		p := packages[i]
		p1 := (p.L + p.W) * 2
		p2 := (p.L + p.H) * 2
		p3 := (p.W + p.H) * 2
		total += p.L*p.W*p.H + min(p1, p2, p3)
	}
	return total
}

func ParsePackages(stream io.Reader) ([]Package, error) {
	var packages []Package
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		dims := bytes.Split([]byte(line), []byte{'x'})
		if len(dims) != 3 {
			return nil, fmt.Errorf("unexpected dimensions: %v", dims)
		}
		l, err := strconv.Atoi(string(dims[0]))
		if err != nil {
			return nil, fmt.Errorf("error parsing length: %w", err)
		}
		w, err := strconv.Atoi(string(dims[1]))
		if err != nil {
			return nil, fmt.Errorf("error parsing width: %w", err)
		}
		h, err := strconv.Atoi(string(dims[2]))
		if err != nil {
			return nil, fmt.Errorf("error parsing height: %w", err)
		}
		packages = append(packages, Package{
			L: l,
			W: w,
			H: h,
		})
	}
	return packages, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	packages, err := ParsePackages(file)
	if err != nil {
		log.Fatalf("error parsing packages: %s", err)
	}
	fmt.Println(TotalRibbon(packages))
}
