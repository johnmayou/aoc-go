package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSecretKey(t *testing.T) {
	cases := []struct {
		prefix string
		want   int
	}{
		{
			prefix: "abcdef",
			want:   609043,
		},
		{
			prefix: "pqrstuv",
			want:   1048970,
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.want, FindSecretKey(tc.prefix))
	}
}
