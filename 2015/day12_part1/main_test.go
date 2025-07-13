package main

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddAllNumbers(t *testing.T) {
	tests := []struct {
		rawJson string
		want    int
	}{
		{rawJson: `[1,2,3]`, want: 6},
		{rawJson: `{"a":2,"b":4}`, want: 6},
		{rawJson: `[[[3]]]`, want: 3},
		{rawJson: `{"a":{"b":4},"c":-1}`, want: 3},
		{rawJson: `{"a":[-1,1]}`, want: 0},
		{rawJson: `[-1,{"a":1}]`, want: 0},
		{rawJson: `[]`, want: 0},
		{rawJson: `{}`, want: 0},
	}

	for i, tt := range tests {
		t.Run("test #"+strconv.Itoa(i), func(t *testing.T) {
			var data any
			require.NoError(t, json.Unmarshal([]byte(tt.rawJson), &data))
			require.Equal(t, tt.want, AddAllNumbers(data))
		})
	}
}
