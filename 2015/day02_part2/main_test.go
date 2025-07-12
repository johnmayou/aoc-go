package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTotalRibbon(t *testing.T) {
	cases := []struct {
		packages []Package
		want     int
	}{
		{
			packages: []Package{{L: 2, W: 3, H: 4}},
			want:     34,
		},
		{
			packages: []Package{{L: 1, W: 1, H: 10}},
			want:     14,
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.want, TotalRibbon(tc.packages))
	}
}

func TestParsePackages(t *testing.T) {
	packages, err := ParsePackages(strings.NewReader("1x2x3\n11x12x13"))
	require.NoError(t, err)
	require.Equal(t, []Package{{L: 1, W: 2, H: 3}, {L: 11, W: 12, H: 13}}, packages)
}
