package main

import "fmt"

func CountCombinations(sizes []int, targetLiters int) int {
	count := 0
	var dfs func(i int, liters int)
	dfs = func(i int, liters int) {
		if liters > targetLiters {
			return
		}
		if liters == targetLiters {
			count += 1
			return
		}
		for j := i; j < len(sizes); j++ {
			dfs(j+1, liters+sizes[j])
		}
	}
	dfs(0, 0)
	return count
}

func main() {
	input := []int{43, 3, 4, 10, 21, 44, 4, 6, 47, 41, 34, 17, 17, 44, 36, 31, 46, 9, 27, 38}
	fmt.Println(CountCombinations(input, 150))
}
