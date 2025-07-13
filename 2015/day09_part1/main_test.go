package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindShortestRoute(t *testing.T) {
	distances := []Distance{
		{From: "London", To: "Dublin", Len: 464},
		{From: "London", To: "Belfast", Len: 518},
		{From: "Dublin", To: "Belfast", Len: 141},
	}
	require.Equal(t, 605, FindShortestRoute(distances))
}

func TestParseDistances(t *testing.T) {
	got, err := ParseDistances(strings.NewReader("London to Dublin = 464"))
	want := []Distance{{From: "London", To: "Dublin", Len: 464}}
	require.NoError(t, err)
	require.Equal(t, want, got)
}
