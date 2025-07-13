package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func LookAndSay(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	var sb strings.Builder
	sb.Grow(len(str) * 2) // help avoid reallocations
	last := str[0]
	lastCnt := 1
	writeLast := func() error {
		if _, err := sb.WriteString(strconv.Itoa(lastCnt)); err != nil {
			return fmt.Errorf("error writing byte: %s", err)
		}
		if err := sb.WriteByte(last); err != nil {
			return fmt.Errorf("error writing byte: %s", err)
		}
		return nil
	}
	for i := 1; i < len(str); i++ {
		if str[i] == last {
			lastCnt++
		} else {
			if err := writeLast(); err != nil {
				return "", err
			}
			last = str[i]
			lastCnt = 1
		}
	}
	if err := writeLast(); err != nil {
		return "", err
	}
	return sb.String(), nil
}

func main() {
	var err error
	str := "3113322113"
	for range 50 {
		str, err = LookAndSay(str)
		if err != nil {
			log.Fatalf("error with look and say: %s", err)
		}
	}
	println(len(str))
}
