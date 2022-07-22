package _0682

import "strconv"

func calPoints(ops []string) int {
	var res int = 0
	arr := make([]int, 0)
	for i := 0; i < len(ops); i++ {
		if ops[i] == "C" {
			arr = arr[:len(arr)-1]
		} else if ops[i] == "+" {
			arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
		} else if ops[i] == "D" {
			arr = append(arr, arr[len(arr)-1]*2)
		} else {
			num, _ := strconv.Atoi(ops[i])
			arr = append(arr, num)
		}
	}

	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}

	return res
}
