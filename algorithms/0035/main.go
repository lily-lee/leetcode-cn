package _0035

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}

	from, end := 0, len(nums)-1
	for from < end {
		mid := (end + from + 1) / 2
		if mid == from || mid == end {
			break
		}
		if target > nums[mid] {
			from = mid
		} else if target == nums[mid] {
			return mid
		} else {
			end = mid
		}
	}
	return from + 1
}
