package _0009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	n := 0
	for y > 0 {
		n = n*10 + y%10
		y = y / 10
	}
	return x == n
}
