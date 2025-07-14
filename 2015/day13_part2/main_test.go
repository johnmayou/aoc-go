package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindOptimalSeating(t *testing.T) {
	scores, err := ParseHappinessScores(strings.NewReader(`Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`))
	require.NoError(t, err)
	_, err = FindOptimalSeating(scores)
	require.NoError(t, err)
}

func TestParseHappinessScores(t *testing.T) {
	tests := map[string]struct {
		in   io.Reader
		want []Happiness
	}{
		"gain": {
			in:   strings.NewReader(`Alice would gain 54 happiness units by sitting next to Bob.`),
			want: []Happiness{{Person: "Alice", Neighbor: "Bob", Happiness: 54}},
		},
		"lose": {
			in:   strings.NewReader(`Alice would lose 79 happiness units by sitting next to Carol.`),
			want: []Happiness{{Person: "Alice", Neighbor: "Carol", Happiness: -79}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseHappinessScores(tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
