package _0746

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	cost = append(cost, 0)
	for i := 2; i < len(cost); i++ {
		cost[i] = min(cost[i-1]+cost[i], cost[i-2]+cost[i])
	}
	return cost[len(cost)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
