package lc_merge

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	sort([][]int{[]int{2, 4}, []int{1, 8}, []int{9, 111}, []int{6, 10}})
}

func TestMerge(t *testing.T) {
	//fmt.Println(merge([][]int{[]int{2, 4}, []int{1, 8}, []int{9, 111}, []int{6, 10}}))
	//fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}))
	//fmt.Println(merge([][]int{[]int{1, 4}, []int{2, 3}}))
	//fmt.Println(merge([][]int{[]int{1, 4}, []int{0, 2}, []int{3, 5}}))
	fmt.Println(merge([][]int{[]int{2, 3}, []int{2, 2}, []int{3, 3}, []int{1, 3}, []int{5, 7}, []int{2, 2}, []int{4, 6}}))
	//fmt.Println(merge([][]int{[]int{0, 4}, []int{3, 5}}))
	//fmt.Println(merge([][]int{[]int{1, 3}, []int{4, 6}, []int{5, 7}}))
}
