package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindNextPass(t *testing.T) {
	tests := []struct {
		pass string
		want string
	}{
		{pass: "abcdefgh", want: "abcdffaa"},
		{pass: "ghijklmn", want: "ghjaabcc"},
	}

	for i, tt := range tests {
		t.Run("test #"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, tt.want, FindNextPass(tt.pass), "input %s, wanted %s", tt.pass, tt.want)
		})
	}
}

func TestIsValidPass(t *testing.T) {
	tests := map[string]struct {
		pass []byte
		want bool
	}{
		"success: #1": {
			pass: []byte("abcdffaa"),
			want: true,
		},
		"success: #2": {
			pass: []byte("ghjaabcc"),
			want: true,
		},
		"fail: contains i and l": {
			pass: []byte("hijklmmn"),
			want: false,
		},
		"fail: no increasing straight of 3 chars": {
			pass: []byte("abbceffg"),
			want: false,
		},
		"fail: only a single char repeat": {
			pass: []byte("abbcegjk"),
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tt.want, IsValidPass(tt.pass), "input %s, wanted %s", tt.pass, tt.want)
		})
	}
}
