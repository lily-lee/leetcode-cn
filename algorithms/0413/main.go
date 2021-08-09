package _0413

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	d, t := nums[1]-nums[0], 0
	total := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == d {
			t++
		} else {
			d, t = nums[i]-nums[i-1], 0
		}
		total += t
	}
	return total
}
