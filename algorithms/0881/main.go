package _0881

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	start, end := 0, len(people)-1
	total := 0
	for start <= end {
		if people[start]+people[end] > limit {
			end--
		} else {
			start++
			end--
		}
		total++
	}
	return total
}
