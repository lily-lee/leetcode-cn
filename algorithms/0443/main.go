package _0443

import "fmt"

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	s := make([]byte, 0)
	j := 1
	for i := 0; i < len(chars); i++ {
		if i < len(chars)-1 && chars[i] == chars[i+1] {
			j++
		} else {
			s = append(s, chars[i])
			chars[len(s)-1] = chars[i]
			if j >= 10 {
				s = append(s, []byte(fmt.Sprintf("%d", j))...)
				chars[len(s)-1] = s[len(s)-1]
				chars[len(s)-2] = s[len(s)-2]
			} else if j > 1 {
				s = append(s, []byte(fmt.Sprintf("%d", j))...)
				chars[len(s)-1] = s[len(s)-1]
			}
			j = 1
		}
	}
	chars = chars[:len(s)]
	return len(s)
}
