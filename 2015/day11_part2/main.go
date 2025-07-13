package main

func FindNextPass(passStr string) string {
	pass := []byte(passStr)
	rotatePass(pass, len(pass)-1)
	for !IsValidPass(pass) {
		rotatePass(pass, len(pass)-1)
	}
	return string(pass)
}

func rotatePass(pass []byte, at int) {
	if pass[at] == 'z' {
		if at > 0 {
			rotatePass(pass, at-1)
		}
		pass[at] = 'a'
	} else {
		pass[at] = byte(int(pass[at]) + 1)
	}
}

func IsValidPass(pass []byte) bool {
	pairStart := -1
	pairRepeat := false
	incStraight := false
	for i := range len(pass) {
		b := pass[i]
		if b == 'i' || b == 'o' || b == 'l' {
			return false
		}
		if !pairRepeat && i+1 < len(pass) && b == pass[i+1] {
			if pairStart == -1 {
				pairStart = i
			} else if pairStart < i-1 {
				pairRepeat = true
			}
		}
		if !incStraight && i+2 < len(pass) {
			bint := uint8(b)
			if pass[i+1] == byte(bint+1) && pass[i+2] == byte(bint+2) {
				incStraight = true
			}
		}
	}
	return pairRepeat && incStraight
}

func main() {
	println(FindNextPass("hxbxxyzz"))
}
