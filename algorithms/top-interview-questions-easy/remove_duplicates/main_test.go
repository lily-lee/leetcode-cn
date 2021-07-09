package remove_duplicates

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a[:2])
	fmt.Println(a[3:])
	fmt.Println(append(a[:2], a[3:]...))
}
