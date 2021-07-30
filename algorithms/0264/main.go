package _0264

func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[i2]*2, dp[i3]*3), dp[i5]*5)
		if dp[i]%2 == 0 {
			i2++
		}
		if dp[i]%3 == 0 {
			i3++
		}
		if dp[i]%5 == 0 {
			i5++
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
