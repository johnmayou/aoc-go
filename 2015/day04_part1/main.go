package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func FindSecretKey(prefix string) int {
	key := 0
	for {
		hash := md5.Sum([]byte(prefix + strconv.Itoa(key)))
		hashStr := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hashStr, "00000") {
			break
		}
		key += 1
	}
	return key
}

func main() {
	fmt.Println(FindSecretKey("yzbqklnj"))
}
