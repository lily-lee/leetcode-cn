package _0001

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if val, ok := m[target-nums[i]]; ok && val != i {
			return []int{i, val}
		}
	}
	return []int{}
}
