package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSues(t *testing.T) {
	got, err := ParseSues(strings.NewReader("Sue 10: children: 1, cats: 2, samoyeds: 3, pomeranians: 4, akitas: 5, vizslas: 6, goldfish: 7, trees: 8, cars: 9, perfumes: 10"))
	want := []Sue{{
		Number: 10,
		Attributes: map[string]uint8{
			"children":    1,
			"cats":        2,
			"samoyeds":    3,
			"pomeranians": 4,
			"akitas":      5,
			"vizslas":     6,
			"goldfish":    7,
			"trees":       8,
			"cars":        9,
			"perfumes":    10,
		},
	}}
	require.NoError(t, err)
	require.Equal(t, want, got)
}
