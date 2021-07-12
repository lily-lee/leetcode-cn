package _0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := merge(nums1, nums2)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return float64(nums[0])
	}
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return float64(nums[mid]+nums[mid-1]) / 2
	}
	return float64(nums[mid])
}

func merge(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	m, n := 0, 0
	tmp := make([]int, 0)
	for m < len(nums1) && n < len(nums2) {
		if nums1[m] < nums2[n] {
			tmp = append(tmp, nums1[m])
			m++
		} else if nums1[m] == nums2[n] {
			tmp = append(tmp, nums1[m], nums2[n])
			m++
			n++
		} else {
			tmp = append(tmp, nums2[n])
			n++
		}
	}
	if m < len(nums1) {
		tmp = append(tmp, nums1[m:]...)
	}
	if n < len(nums2) {
		tmp = append(tmp, nums2[n:]...)
	}
	return tmp
}
