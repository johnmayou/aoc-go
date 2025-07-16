package main

import (
	"fmt"
	"math"
)

func CountMinContainerCombinations(sizes []int, targetLiters int) int {
	containerCount := int(math.Inf(1))
	containerCombinations := 0
	var dfs func(i int, liters int, containers int)
	dfs = func(i int, liters int, containers int) {
		if liters > targetLiters {
			return
		}
		if liters == targetLiters {
			if containers < containerCount {
				containerCount = containers
				containerCombinations = 1
			} else if containers == containerCount {
				containerCombinations += 1
			}
			return
		}
		for j := i; j < len(sizes); j++ {
			dfs(j+1, liters+sizes[j], containers+1)
		}
	}
	dfs(0, 0, 0)
	return containerCombinations
}

func main() {
	input := []int{43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38}
	fmt.Println(CountMinContainerCombinations(input, 150))
}
