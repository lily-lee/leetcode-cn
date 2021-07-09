package _1049

import (
	"fmt"
	"testing"
)

func TestLastStoneWeightII(t *testing.T) {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func TestSort(t *testing.T) {
	stones := []int{2, 7, 4, 1, 8, 1}
	sort(stones)
	fmt.Println(stones)
}

func TestSliceHandle(t *testing.T) {
	sliceHandle([]int{2, 7, 4, 1, 8, 1}, 1)
}
