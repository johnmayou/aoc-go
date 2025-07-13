package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNice(t *testing.T) {
	cases := map[string]struct {
		str  string
		want bool
	}{
		"nice #1": {
			str:  "qjhvhtzxzqqjkmpb",
			want: true,
		},
		"nice #2": {
			str:  "xxyxx",
			want: true,
		},
		"not nice: no char repeat with single letter between them": {
			str:  "uurcxstgmygtbstg",
			want: false,
		},
		"not nice: no pair that appears twice": {
			str:  "ieodomkazucvgmuy",
			want: false,
		},
	}

	for tcName, tc := range cases {
		t.Run(tcName, func(t *testing.T) {
			assert.Equal(t, tc.want, IsNice(tc.str))
		})
	}
}
