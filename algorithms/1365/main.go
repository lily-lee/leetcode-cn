package _1365

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	type pair = struct {
		index int
		value int
	}
	m := make([]pair, len(nums))
	for i := 0; i < len(nums); i++ {
		m[i] = pair{i, nums[i]}
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].value < m[j].value
	})
	result := make([]int, len(nums))
	p := -1
	for i, v := range m {
		if p == -1 || m[i-1].value != m[i].value {
			p = i
		}
		result[v.index] = p
	}
	return result
}
