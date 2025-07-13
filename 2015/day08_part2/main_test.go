package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringEncodingLengthDiff(t *testing.T) {
	tests := []struct {
		str  string
		want int
	}{
		{
			str:  `""`,
			want: 4,
		},
		{
			str:  `"abc"`,
			want: 4,
		},
		{
			str:  `"aaa\"aaa"`,
			want: 6,
		},
		{
			str:  `"\x27"`,
			want: 5,
		},
	}

	for i, tt := range tests {
		t.Run("test #"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, tt.want, StringEncodingLengthDiff(tt.str))
		})
	}
}
