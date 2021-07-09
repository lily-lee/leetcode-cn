package offer61

import (
	"fmt"
	"testing"
)

func TestIsStraight(t *testing.T) {
	fmt.Println(isStraight([]int{0, 0, 1, 2, 3}))
	fmt.Println(isStraight([]int{0, 1, 0, 2, 3}))
	fmt.Println(isStraight([]int{1, 0, 0, 3, 4}))
	fmt.Println(isStraight([]int{0, 1, 0, 3, 4}))
	fmt.Println(isStraight([]int{0, 0, 1, 3, 5}))
	fmt.Println(isStraight([]int{0, 0, 1, 4, 6}))
}

func TestSort(t *testing.T) {
	sort([]int{5, 6, 1, 5, 3, 0})
}
