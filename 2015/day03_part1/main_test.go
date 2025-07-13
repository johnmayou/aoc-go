package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDeliveries(t *testing.T) {
	cases := []struct {
		stream io.Reader
		want   int
	}{
		{
			stream: strings.NewReader(">"),
			want:   2,
		},
		{
			stream: strings.NewReader("^>v<"),
			want:   4,
		},
		{
			stream: strings.NewReader("^v^v^v^v^v"),
			want:   2,
		},
	}

	for _, tc := range cases {
		got, err := CountUniqueDeliveries(tc.stream)
		assert.NoError(t, err)
		assert.Equal(t, tc.want, got)
	}
}
