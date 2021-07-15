package _0007

import (
	"math"
)

func reverse(x int) int {
	rev := 0
	for x != 0 {
		rev = rev*10 + x%10
		x = x / 10
		if rev < math.MinInt32 || rev > math.MaxInt32 {
			return 0
		}
	}
	return rev
}
