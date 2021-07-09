package _3

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if len(nums) == 1 {
		return -1
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	if len(nums) == 2 {
		return -1
	}
	index := getTarget(nums, target, 0, len(nums)-1)

	return index
}

func getTarget(nums []int, target, start, end int) int {
	mid := getMid(start, end)

	if target == nums[start] {
		return start
	}
	if target == nums[mid] {
		return mid
	}
	if target == nums[end] {
		return end
	}
	if end == mid && target != nums[mid] {
		return -1
	}
	if nums[start] > nums[end] {
		fmt.Println(nums[start], nums[end])
		if nums[end] > target && target > nums[mid] {
			return getTarget(nums, target, mid, end)
		}
		if target > nums[mid] && nums[mid] > nums[start] {
			return getTarget(nums, target, mid, end)
		}
		if nums[mid] > nums[start] && nums[end] > target {
			return getTarget(nums, target, mid, end)
		}

		if nums[end] > nums[mid] && nums[mid] > target {
			return getTarget(nums, target, start, mid)
		}
		if nums[mid] > target && target > nums[start] {
			return getTarget(nums, target, start, mid)
		}
		if target > nums[start] && nums[end] > nums[mid] {
			return getTarget(nums, target, start, mid)
		}
	} else if nums[start] < nums[end] {
		if target > nums[mid] {
			return getTarget(nums, target, mid, end)
		} else if target == nums[mid] {
			return mid
		} else {
			return getTarget(nums, target, start, mid)
		}
	}

	return -1
}

func getMid(start, end int) int {
	total := start + end
	if total%2 == 0 {
		return total / 2
	}
	return total/2 + 1
}
