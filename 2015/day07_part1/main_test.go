package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssemble(t *testing.T) {
	instructions, err := ParseInstructions(
		strings.NewReader(`123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`),
	)
	require.NoError(t, err)
	want := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}
	got, err := Assemble(instructions)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestParseInstructions(t *testing.T) {
	tests := map[string]struct {
		stream io.Reader
		want   []Instruction
	}{
		"direct": {
			stream: strings.NewReader("123 -> x"),
			want:   []Instruction{DirectValueInstruction{Value: Operand{IsLiteral: true, Value: 123}, DestSignal: "x"}},
		},
		"bitwise and": {
			stream: strings.NewReader("x AND y -> d"),
			want:   []Instruction{BitwiseAndInstruction{Left: Operand{IsLiteral: false, Signal: "x"}, Right: Operand{IsLiteral: false, Signal: "y"}, DestSignal: "d"}},
		},
		"bitwise or": {
			stream: strings.NewReader("x OR y -> e"),
			want:   []Instruction{BitwiseOrInstruction{Left: Operand{IsLiteral: false, Signal: "x"}, Right: Operand{IsLiteral: false, Signal: "y"}, DestSignal: "e"}},
		},
		"bitwise lshift": {
			stream: strings.NewReader("x LSHIFT 2 -> f"),
			want:   []Instruction{BitwiseLShiftInstruction{Left: Operand{IsLiteral: false, Signal: "x"}, Right: Operand{IsLiteral: true, Value: 2}, DestSignal: "f"}},
		},
		"bitwise rshift": {
			stream: strings.NewReader("y RSHIFT 2 -> g"),
			want:   []Instruction{BitwiseRShiftInstruction{Left: Operand{IsLiteral: false, Signal: "y"}, Right: Operand{IsLiteral: true, Value: 2}, DestSignal: "g"}},
		},
		"bitwise not": {
			stream: strings.NewReader("NOT y -> i"),
			want:   []Instruction{BitwiseNotInstruction{Value: Operand{IsLiteral: false, Signal: "y"}, DestSignal: "i"}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseInstructions(tt.stream)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
