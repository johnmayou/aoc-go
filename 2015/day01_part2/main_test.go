package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindFirstBasementPos(t *testing.T) {
	cases := []struct {
		stream io.Reader
		want   int
	}{
		{
			stream: strings.NewReader(")"),
			want:   1,
		},
		{
			stream: strings.NewReader("()())"),
			want:   5,
		},
	}

	for _, tc := range cases {
		got, err := FindFirstBasementPos(tc.stream)
		require.NoError(t, err)
		require.Equal(t, tc.want, got)
	}
}
