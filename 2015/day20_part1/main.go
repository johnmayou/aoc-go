package main

import (
	"fmt"
)

func FindLowestHouseNumber(targetPresents int, upperBound int) int {
	presents := make([]int, upperBound+1)
	for elf := 1; elf <= upperBound; elf++ {
		house := elf
		for house <= upperBound {
			presents[house] += elf
			house += elf
		}
	}
	for house, total := range presents {
		if total*10 >= targetPresents {
			return house
		}
	}
	return -1
}

func main() {
	fmt.Println(FindLowestHouseNumber(33_100_000, 1_000_000))
}
