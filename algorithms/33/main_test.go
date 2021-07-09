package _3

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		if r := search([]int{}, 1); r != -1 {
			t.Fail()
		}
	})
	t.Run("1-1", func(t *testing.T) {
		if r := search([]int{1}, 1); r != 0 {
			t.Fail()
		}
	})
	t.Run("1-2", func(t *testing.T) {
		if r := search([]int{1}, 0); r != -1 {
			t.Fail()
		}
	})
	t.Run("2-1", func(t *testing.T) {
		if r := search([]int{1, 3}, 0); r != -1 {
			fmt.Println("r:", r)
			t.Fail()
		}
	})
	t.Run("2-2", func(t *testing.T) {
		if r := search([]int{1, 3}, 1); r != 0 {
			fmt.Println("r:", r)
			t.Fail()
		}
	})
	t.Run("2-3", func(t *testing.T) {
		if r := search([]int{1, 3}, 3); r != 1 {
			fmt.Println("r:", r)
			t.Fail()
		}
	})
	t.Run("3-1", func(t *testing.T) {
		if r := search([]int{3, 0, 1}, 10); r != -1 {
			t.Fail()
		}
	})
	t.Run("3-2", func(t *testing.T) {
		if r := search([]int{3, 0, 1}, 0); r != 1 {
			t.Fail()
		}
	})
	t.Run("3-3", func(t *testing.T) {
		if r := search([]int{2, 3, 1}, 0); r != -1 {
			t.Fail()
		}
	})
	t.Run("3-4", func(t *testing.T) {
		if r := search([]int{1, 3, 5}, 0); r != -1 {
			t.Fail()
		}
	})
	t.Run("3-4", func(t *testing.T) {
		if r := search([]int{2, 3, 1}, 3); r != 1 {
			t.Fail()
		}
	})
	t.Run("3-5", func(t *testing.T) {
		if r := search([]int{1, 2, 3}, 1); r != 0 {
			t.Fail()
		}
	})
	t.Run("4-1", func(t *testing.T) {
		if r := search([]int{1, 2, 3, 4}, 5); r != -1 {
			t.Fail()
		}
	})
	t.Run("4-2", func(t *testing.T) {
		if r := search([]int{1, 2, 3, 4}, 0); r != -1 {
			t.Fail()
		}
	})
	t.Run("4-3", func(t *testing.T) {
		if r := search([]int{1, 2, 3, 4}, 3); r != 2 {
			t.Fail()
		}
	})
	t.Run("4-4", func(t *testing.T) {
		if r := search([]int{1, 2, 3, 4}, 2); r != 1 {
			t.Fail()
		}
	})
	t.Run("4-5", func(t *testing.T) {
		if r := search([]int{2, 3, 4, 1}, 5); r != -1 {
			t.Fail()
		}
	})
	t.Run("4-6", func(t *testing.T) {
		if r := search([]int{2, 3, 4, 1}, 3); r != 1 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4); r != 3 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 8); r != 4 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 4); r != 0 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 5); r != 1 {
			fmt.Println(r)
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 1); r != 6 {
			fmt.Println(r)
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8); r != 4 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 9, 1, 2}, 10); r != -1 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 9); r != -1 {
			fmt.Println("...", r)
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 0, 1, 2}, 6); r != 2 {
			fmt.Println("here", r)
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{4, 5, 6, 7, 0, 1, 2}, 3); r != -1 {
			fmt.Println(r)
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{10, 11, 12, 13, 14, 1, 2, 3, 4, 5, 6, 7, 8}, 9); r != -1 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7}, 8); r != -1 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7}, 0); r != 5 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7}, 11); r != 1 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7}, 12); r != 2 {
			t.Fail()
		}
	})
	t.Run("", func(t *testing.T) {
		if r := search([]int{10, 11, 12, 13, 14, 0, 1, 2, 3, 4, 5, 6, 7}, 13); r != 3 {
			t.Fail()
		}
	})
}
