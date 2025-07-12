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

func TotalWrapping(packages []Package) int {
	total := 0
	for i := range packages {
		p := packages[i]
		s1 := p.L * p.W
		s2 := p.L * p.H
		s3 := p.W * p.H
		total += s1*2 + s2*2 + s3*2 + min(s1, s2, s3)
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
	fmt.Println(TotalWrapping(packages))
}
