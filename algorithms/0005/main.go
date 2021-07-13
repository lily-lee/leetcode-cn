package _0005

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	str := s[0:1]
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return str
	}
	for i := 0; i < len(s); i++ {
		j := 2
		for j <= len(s) && i <= len(s)-j {
			sub := s[i : i+j]
			if check(sub) {
				if len(sub) > len(str) {
					str = sub
				}
			}
			j++
		}
	}

	return str
}

func check(s string) bool {
	mid := len(s) / 2
	if len(s)%2 == 0 {
		mid = mid - 1
	}
	for i := 0; i <= mid; i++ {
		if !(s[i] == s[len(s)-i-1]) {
			return false
		}
	}
	return true
}
