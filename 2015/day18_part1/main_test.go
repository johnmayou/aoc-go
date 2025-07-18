package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestGridStep(t *testing.T) {
	grid, err := ParseGrid(strings.NewReader(`.#.#.#
...##.
#....#
..#...
#.#..#
####..`), 6, 6)
	require.NoError(t, err)
	require.NoError(t, GridStep(grid))
	want, err := ParseGrid(strings.NewReader(`..##..
..##.#
...##.
......
#.....
#.##..`), 6, 6)
	require.NoError(t, err)
	if diff := cmp.Diff(want, grid); diff != "" {
		t.Errorf("grid mismatch (-want +got):\n%s", diff)
	}
}

func TestParseGrid(t *testing.T) {
	got, err := ParseGrid(strings.NewReader("#.\n.#"), 2, 2)
	want := [][]uint8{{StateOn, StateOff}, {StateOff, StateOn}}
	require.NoError(t, err)
	require.Equal(t, want, got)
}
