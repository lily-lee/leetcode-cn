package _0008

import (
	"math"
)

func myAtoi(s string) int {
	num := make([]byte, 0)
	flag := 1
	for i := 0; i < len(s); i++ {
		if !isNum(s[i]) && len(num) != 0 {
			break
		}
		if !isNum(s[i]) && len(num) == 0 {
			if isSpace(s[i]) {
				continue
			}
			if !isSig(s[i]) {
				break
			}
		}
		if i < len(s)-1 && isSig(s[i]) && !isNum(s[i+1]) {
			break
		}
		if isSig(s[i]) && s[i] == '-' {
			flag = -1
			continue
		}
		if isNum(s[i]) {
			num = append(num, s[i])
		}
	}
	if len(num) == 0 {
		return 0
	}
	n := 0
	for i := 0; i < len(num); i++ {
		if n*flag < math.MinInt32 {
			return math.MinInt32
		}
		if n*flag > math.MaxInt32 {
			return math.MaxInt32
		}
		n = n*10 + int(num[i]-'0')
	}
	n = n * flag
	if n < math.MinInt32 {
		return math.MinInt32
	}
	if n > math.MaxInt32 {
		return math.MaxInt32
	}
	return n
}
func isSpace(b byte) bool {
	if b == ' ' {
		return true
	}
	return false
}
func isNum(b byte) bool {
	if b < '0' || b > '9' {
		return false
	}
	return true
}
func isSig(b byte) bool {
	if b == '+' || b == '-' {
		return true
	}
	return false
}
