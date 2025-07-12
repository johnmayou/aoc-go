package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindFloor(t *testing.T) {
	cases := []struct {
		stream io.Reader
		want   int
	}{
		{
			stream: strings.NewReader("(())"),
			want:   0,
		},
		{
			stream: strings.NewReader("()()"),
			want:   0,
		},
		{
			stream: strings.NewReader("))((((("),
			want:   3,
		},
	}

	for _, tc := range cases {
		got, err := FindFloor(tc.stream)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}
