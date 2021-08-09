package offerII006

func twoSum(numbers []int, target int) []int {
	x, y := 0, len(numbers)-1
	if numbers[len(numbers)/2] > target {
		y = len(numbers) / 2
	}
	for x < y {
		if numbers[x]+numbers[y] == target {
			break
		} else if numbers[x]+numbers[y] > target {
			y--
		} else {
			x++
		}
	}
	return []int{x, y}
}
