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
			str:  "ugknbfddgicrmopn",
			want: true,
		},
		"nice #2": {
			str:  "aaa",
			want: true,
		},
		"not nice: no double letter": {
			str:  "jchzalrnumimnmhp",
			want: false,
		},
		"not nice: contains xy": {
			str:  "haegwjzuvuyypxyu",
			want: false,
		},
		"not nice: contains only one vowel": {
			str:  "dvszwmarrgswjxmb",
			want: false,
		},
	}

	for tcName, tc := range cases {
		t.Run(tcName, func(t *testing.T) {
			assert.Equal(t, tc.want, IsNice(tc.str))
		})
	}
}
