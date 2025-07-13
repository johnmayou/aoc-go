package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringLiteralToInMemoryDiff(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			str:  `""`,
			want: 2,
		},
		{
			str:  `"abc"`,
			want: 2,
		},
		{
			str:  `"aaa\"aaa"`,
			want: 3,
		},
		{
			str:  `"\x27"`,
			want: 5,
		},
	}

	for i, tt := range tests {
		t.Run("test #"+strconv.Itoa(i), func(t *testing.T) {
			got, err := StringLiteralToInMemoryDiff(tt.str)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
