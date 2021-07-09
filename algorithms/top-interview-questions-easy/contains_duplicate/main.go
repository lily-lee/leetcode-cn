package contains_duplicate

func containsDuplicate(nums []int) bool {
	result := false
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	for _, v := range m {
		if v > 1 {
			result = true
		}
	}
	return result
}
