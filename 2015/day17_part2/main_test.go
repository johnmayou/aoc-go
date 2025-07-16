package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountMinContainerCombinations(t *testing.T) {
	require.Equal(t, CountMinContainerCombinations([]int{5, 5, 10, 15, 20}, 25), 3)
}
