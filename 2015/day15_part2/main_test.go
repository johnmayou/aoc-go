package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindBestRecipe(t *testing.T) {
	ingredients := []Ingredient{
		{Name: "Butterscotch", Capacity: -1, Durability: -2, Flavor: 6, Texture: 3, Calories: 8},
		{Name: "Cinnamon", Capacity: 2, Durability: 3, Flavor: -2, Texture: -1, Calories: 3},
	}
	recipe, score := FindBestRecipe(ingredients, 500, 100)
	assert.Equal(t, map[string]int{"Butterscotch": 40, "Cinnamon": 60}, recipe)
	assert.Equal(t, 57600000, score)
}

func TestParseIngredients(t *testing.T) {
	tests := map[string]struct {
		in   io.Reader
		want []Ingredient
	}{
		"positive": {
			in:   strings.NewReader("Name: capacity 10, durability 11, flavor 12, texture 13, calories 14"),
			want: []Ingredient{{Name: "Name", Capacity: 10, Durability: 11, Flavor: 12, Texture: 13, Calories: 14}},
		},
		"negative": {
			in:   strings.NewReader("Name: capacity -10, durability -11, flavor -12, texture -13, calories -14"),
			want: []Ingredient{{Name: "Name", Capacity: -10, Durability: -11, Flavor: -12, Texture: -13, Calories: -14}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseIngredients(tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
