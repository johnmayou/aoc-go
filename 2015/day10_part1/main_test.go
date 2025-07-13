package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookAndSay(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{str: "1", want: "11"},
		{str: "11", want: "21"},
		{str: "21", want: "1211"},
		{str: "1211", want: "111221"},
		{str: "111221", want: "312211"},
	}

	for i, tt := range tests {
		t.Run("test #"+strconv.Itoa(i), func(t *testing.T) {
			require.Equal(t, tt.want, LookAndSay(tt.str), "input '%s' wanted '%s'", tt.str, tt.want)
		})
	}
}
