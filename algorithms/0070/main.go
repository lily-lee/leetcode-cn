package _0070

func climbStairs(n int) int {
	// n = n-1 + 1
	// n = n-2 + 2
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i <= 2 {
			dp[i] = i
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}
