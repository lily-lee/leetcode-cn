package _1049

import "fmt"

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	for n > 1 {
		sort(stones)
		i := stones[0] - stones[1]
		if i > 0 {
			stones[1] = i
			stones = stones[1:]
		} else {
			stones = stones[2:]
		}
		fmt.Println(stones)
		n = len(stones)
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}

func sort(stones []int) {
	for i := 0; i < len(stones); i++ {
		for j := i + 1; j < len(stones); j++ {
			if stones[i] < stones[j] {
				stones[i], stones[j] = stones[j], stones[i]
			}
		}
	}
}

func sliceHandle(stones []int, n int) {
	fmt.Println(stones[n:])
}
