package offer61

import "fmt"

func isStraight(nums []int) bool {
	// 先排序
	sort(nums)

	n, first := 0, 0
	result := true

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			n, first = i, i
			break
		}
	}

	for i := first; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i] - 1
		if diff <= n && diff >= 0 {
			n = n - diff
			continue
		} else {
			result = false
		}
	}

	return result
}

func sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	fmt.Print(nums)
}
