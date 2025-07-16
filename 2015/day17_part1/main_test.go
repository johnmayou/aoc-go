package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountCombinations(t *testing.T) {
	require.Equal(t, CountCombinations([]int{5, 5, 10, 15, 20}, 25), 4)
}
