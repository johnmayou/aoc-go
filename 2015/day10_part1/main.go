package main

import "strconv"

func LookAndSay(str string) string {
	if len(str) == 0 {
		return ""
	}
	var result string
	last := str[0]
	lastCnt := 1
	for i := 1; i < len(str); i++ {
		if str[i] == last {
			lastCnt++
		} else {
			result += strconv.Itoa(lastCnt) + string(last)
			last = str[i]
			lastCnt = 1
		}
	}
	result += strconv.Itoa(lastCnt) + string(last)
	return result
}

func main() {
	str := "3113322113"
	for range 40 {
		str = LookAndSay(str)
	}
	println(len(str))
}
