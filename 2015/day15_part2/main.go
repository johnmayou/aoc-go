package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"regexp"
	"strconv"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func FindBestRecipe(ingredients []Ingredient, totalCalories int, totalTeaspoons int) (map[string]int, int) {
	ingMap := make(map[string]Ingredient, len(ingredients))
	for i := range ingredients {
		ingMap[ingredients[i].Name] = ingredients[i]
	}
	maxRecipe := make(map[string]int, 0)
	maxScore := 0
	var dfs func(i int, calories int, teaspoons int, recipe map[string]int)
	dfs = func(i int, calories int, teaspoons int, recipe map[string]int) {
		if calories > totalCalories {
			return
		}
		if i == len(ingredients)-1 {
			tspsLeft := totalTeaspoons - teaspoons
			ing := ingredients[i]
			recipe[ing.Name] = tspsLeft
			calories += ing.Calories * tspsLeft
			if calories != totalCalories {
				return
			}
			capacity, durability, flavor, texture := 0, 0, 0, 0
			for name, tsps := range recipe {
				ing := ingMap[name]
				capacity += ing.Capacity * tsps
				durability += ing.Durability * tsps
				flavor += ing.Flavor * tsps
				texture += ing.Texture * tsps
			}
			score := max(0, capacity) * max(0, durability) * max(0, flavor) * max(0, texture)
			if score > maxScore {
				maxScore = score
				maxRecipe = maps.Clone(recipe)
			}
			return
		}
		ing := ingredients[i]
		for tsps := range totalTeaspoons - teaspoons + 1 {
			recipe[ing.Name] = tsps
			dfs(i+1, calories+ing.Calories*tsps, teaspoons+tsps, recipe)
		}
	}
	dfs(0, 0, 0, make(map[string]int, len(ingredients)))
	return maxRecipe, maxScore
}

var ingredientRegex = regexp.MustCompile(`^(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)$`)

func ParseIngredients(in io.Reader) ([]Ingredient, error) {
	var ingredients []Ingredient
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		matches := ingredientRegex.FindStringSubmatch(line)
		if matches == nil {
			return nil, fmt.Errorf("invalid ingredient format: %s", matches)
		}
		capacity, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("error converting capacity (%s): %s", matches[2], err)
		}
		durability, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("error converting durability (%s): %s", matches[3], err)
		}
		flavor, err := strconv.Atoi(matches[4])
		if err != nil {
			return nil, fmt.Errorf("error converting flavor (%s): %s", matches[4], err)
		}
		texture, err := strconv.Atoi(matches[5])
		if err != nil {
			return nil, fmt.Errorf("error converting texture (%s): %s", matches[5], err)
		}
		calories, err := strconv.Atoi(matches[6])
		if err != nil {
			return nil, fmt.Errorf("error converting calories (%s): %s", matches[6], err)
		}
		ingredients = append(ingredients, Ingredient{
			Name:       matches[1],
			Capacity:   capacity,
			Durability: durability,
			Flavor:     flavor,
			Texture:    texture,
			Calories:   calories,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning line: %s", err)
	}
	return ingredients, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	ingredients, err := ParseIngredients(file)
	if err != nil {
		log.Fatalf("error parsing ingredients: %s", err)
	}
	recipe, score := FindBestRecipe(ingredients, 500, 100)
	fmt.Printf("recipe with %d score won: %+v\n", score, recipe)
}
