package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRace(t *testing.T) {
	reindeer := []Reindeer{
		{
			Name:        "Comet",
			FlySpeed:    14,
			FlySeconds:  10,
			RestSeconds: 127,
		},
		{
			Name:        "Dancer",
			FlySpeed:    16,
			FlySeconds:  11,
			RestSeconds: 162,
		},
	}
	distance, err := Race(reindeer, 1000)
	require.NoError(t, err)
	assert.Equal(t, 689, distance)
}

func TestParseReindeer(t *testing.T) {
	got, err := ParseReindeer(strings.NewReader("Cupid can fly 22 km/s for 2 seconds, but then must rest for 41 seconds."))
	want := []Reindeer{{
		Name:        "Cupid",
		FlySpeed:    22,
		FlySeconds:  2,
		RestSeconds: 41,
	}}
	require.NoError(t, err)
	require.Equal(t, want, got)
}
