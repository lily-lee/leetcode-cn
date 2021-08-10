package _0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if j >= n {
			sorted = append(sorted, nums1[i:]...)
			break
		}
		if i >= m {
			sorted = append(sorted, nums2[j:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, tail := m-1, n-1, m+n-1
	for {
		if tail == -1 {
			break
		}
		if j == -1 {
			nums1[tail] = nums1[i]
			i--
		} else if i == -1 {
			nums1[tail] = nums2[j]
			j--
		} else if nums1[i] >= nums2[j] {
			nums1[tail] = nums1[i]
			i--
		} else {
			nums1[tail] = nums2[j]
			j--
		}
		tail--
	}
}
