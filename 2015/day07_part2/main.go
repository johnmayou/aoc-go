package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type InstructionType int

const (
	DirectValue InstructionType = iota
	BitwiseAnd
	BitwiseOr
	BitwiseLShift
	BitwiseRShift
	BitwiseNot
)

type Instruction interface {
	Type() InstructionType
}

type Operand struct {
	IsLiteral bool
	Value     uint16 // if IsLiteral
	Signal    string // if !IsLiteral
}

type DirectValueInstruction struct {
	Value      Operand
	DestSignal string
}

func (i DirectValueInstruction) Type() InstructionType { return DirectValue }

type BitwiseAndInstruction struct {
	Left, Right Operand
	DestSignal  string
}

func (i BitwiseAndInstruction) Type() InstructionType { return BitwiseAnd }

type BitwiseOrInstruction struct {
	Left, Right Operand
	DestSignal  string
}

func (i BitwiseOrInstruction) Type() InstructionType { return BitwiseOr }

type BitwiseLShiftInstruction struct {
	Left, Right Operand
	DestSignal  string
}

func (i BitwiseLShiftInstruction) Type() InstructionType { return BitwiseLShift }

type BitwiseRShiftInstruction struct {
	Left, Right Operand
	DestSignal  string
}

func (i BitwiseRShiftInstruction) Type() InstructionType { return BitwiseRShift }

type BitwiseNotInstruction struct {
	Value      Operand
	DestSignal string
}

func (i BitwiseNotInstruction) Type() InstructionType { return BitwiseNot }

func Assemble(instructions []Instruction) (map[string]uint16, error) {
	signals := make(map[string]uint16)
	evalOperand := func(op Operand) (uint16, bool) {
		if op.IsLiteral {
			return op.Value, true
		}
		if op.Signal == "b" {
			return 46065, true
		}
		val, ok := signals[op.Signal]
		return val, ok
	}
	q := list.New()
	for i := range instructions {
		q.PushBack(instructions[i])
	}
	for q.Len() > 0 {
		instruction := q.Remove(q.Front())
		switch instr := instruction.(type) {
		case DirectValueInstruction:
			if val, ok := evalOperand(instr.Value); ok {
				signals[instr.DestSignal] = val
			} else {
				q.PushBack(instr)
			}
		case BitwiseAndInstruction:
			left, okL := evalOperand(instr.Left)
			right, okR := evalOperand(instr.Right)
			if okL && okR {
				signals[instr.DestSignal] = left & right
			} else {
				q.PushBack(instr)
			}
		case BitwiseOrInstruction:
			left, okL := evalOperand(instr.Left)
			right, okR := evalOperand(instr.Right)
			if okL && okR {
				signals[instr.DestSignal] = left | right
			} else {
				q.PushBack(instr)
			}
		case BitwiseLShiftInstruction:
			left, okL := evalOperand(instr.Left)
			right, okR := evalOperand(instr.Right)
			if okL && okR {
				signals[instr.DestSignal] = left << right
			} else {
				q.PushBack(instr)
			}
		case BitwiseRShiftInstruction:
			left, okL := evalOperand(instr.Left)
			right, okR := evalOperand(instr.Right)
			if okL && okR {
				signals[instr.DestSignal] = left >> right
			} else {
				q.PushBack(instr)
			}
		case BitwiseNotInstruction:
			if val, ok := evalOperand(instr.Value); ok {
				signals[instr.DestSignal] = ^val
			} else {
				q.PushBack(instr)
			}
		default:
			return nil, fmt.Errorf("invalid instruction type: %v", instr)
		}
	}
	return signals, nil
}

func ParseInstructions(stream io.Reader) ([]Instruction, error) {
	var instructions []Instruction
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error scanning line: %s", err)
		}
		var instruction Instruction
		if matches := regexp.MustCompile(`^(\w+) -> (\w+)$`).FindStringSubmatch(line); matches != nil {
			instruction = DirectValueInstruction{
				Value:      parseOperand(matches[1]),
				DestSignal: matches[2],
			}
		} else if matches := regexp.MustCompile(`^(\w+) AND (\w+) -> (\w+)$`).FindStringSubmatch(line); matches != nil {
			instruction = BitwiseAndInstruction{
				Left:       parseOperand(matches[1]),
				Right:      parseOperand(matches[2]),
				DestSignal: matches[3],
			}
		} else if matches := regexp.MustCompile(`^(\w+) OR (\w+) -> (\w+)$`).FindStringSubmatch(line); matches != nil {
			instruction = BitwiseOrInstruction{
				Left:       parseOperand(matches[1]),
				Right:      parseOperand(matches[2]),
				DestSignal: matches[3],
			}
		} else if matches := regexp.MustCompile(`^(\w+) LSHIFT (\w+) -> (\w+)$`).FindStringSubmatch(line); matches != nil {
			instruction = BitwiseLShiftInstruction{
				Left:       parseOperand(matches[1]),
				Right:      parseOperand(matches[2]),
				DestSignal: matches[3],
			}
		} else if matches := regexp.MustCompile(`^(\w+) RSHIFT (\w+) -> (\w+)$`).FindStringSubmatch(line); matches != nil {
			instruction = BitwiseRShiftInstruction{
				Left:       parseOperand(matches[1]),
				Right:      parseOperand(matches[2]),
				DestSignal: matches[3],
			}
		} else if matches := regexp.MustCompile(`^NOT (\w+) -> (\w+)$`).FindStringSubmatch(line); matches != nil {
			instruction = BitwiseNotInstruction{
				Value:      parseOperand(matches[1]),
				DestSignal: matches[2],
			}
		} else {
			return nil, fmt.Errorf("invalid instruction: %s", line)
		}
		instructions = append(instructions, instruction)
	}
	return instructions, nil
}

func parseOperand(token string) Operand {
	if val, err := strconv.Atoi(token); err == nil {
		return Operand{IsLiteral: true, Value: uint16(val)}
	}
	return Operand{IsLiteral: false, Signal: token}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file: %s", err)
	}
	defer file.Close()
	instructions, err := ParseInstructions(file)
	if err != nil {
		log.Fatalf("error parsing instructions: %s", err)
	}
	signals, err := Assemble(instructions)
	if err != nil {
		log.Fatalf("error assembling signals: %s", err)
	}
	if val, ok := signals["a"]; ok {
		println(val)
	} else {
		log.Fatalf("no signal for wire a: %v", signals)
	}
}
