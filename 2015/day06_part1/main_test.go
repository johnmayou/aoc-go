package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerform(t *testing.T) {
	instructions := []Instruction{
		{Action: TurnOn, Start: Coord{X: 0, Y: 0}, Stop: Coord{X: 0, Y: 1}},  // on: (0, 0), (0, 1)
		{Action: Toggle, Start: Coord{X: 0, Y: 0}, Stop: Coord{X: 1, Y: 0}},  // on: (1, 0), (0, 1)
		{Action: TurnOff, Start: Coord{X: 0, Y: 1}, Stop: Coord{X: 0, Y: 1}}, // on: (1, 0)
	}
	grid := make([][]bool, 2)
	for i := range grid {
		grid[i] = make([]bool, 2)
	}
	Perform(grid, instructions)
	require.Equal(t, 1, CountOn(grid))
}

func TestCountOn(t *testing.T) {
	require.Equal(t, 3, CountOn([][]bool{{false, true}, {true, true}}))
}

func TestParseInstructions(t *testing.T) {
	cases := map[string]struct {
		stream io.Reader
		want   []Instruction
	}{
		"turn on": {
			stream: strings.NewReader("turn on 0,0 through 999,999"),
			want:   []Instruction{{Action: TurnOn, Start: Coord{X: 0, Y: 0}, Stop: Coord{X: 999, Y: 999}}},
		},
		"turn off": {
			stream: strings.NewReader("turn off 499,499 through 500,500"),
			want:   []Instruction{{Action: TurnOff, Start: Coord{X: 499, Y: 499}, Stop: Coord{X: 500, Y: 500}}},
		},
		"toggle": {
			stream: strings.NewReader("toggle 0,0 through 999,0"),
			want:   []Instruction{{Action: Toggle, Start: Coord{X: 0, Y: 0}, Stop: Coord{X: 999, Y: 0}}},
		},
	}

	for tcName, tc := range cases {
		t.Run(tcName, func(t *testing.T) {
			got, err := ParseInstructions(tc.stream)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
