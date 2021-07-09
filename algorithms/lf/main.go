package lf

func search(nums []int, target int) int {
	start := 0
	end := len(nums)

	for start < end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			if nums[end] > target {
				start = middle
				continue
			}
			return middle + search(nums[middle:], target)
		}

		if nums[middle] > target {
			if nums[start] < target {
				end = middle
				continue
			}
			return start + search(nums[:middle], target)
		}
	}
}
