package two_sum

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[target-nums[i]] = i
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		index, ok := m[nums[i]]
		if !ok {
			continue
		}
		if index == i {
			continue
		}
		result = []int{index, i}
	}

	return result
}
