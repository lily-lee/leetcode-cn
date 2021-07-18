package _0011

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	x1, x2 := 0, len(height)-1
	w := x2 - x1
	h := min(height[x2], height[x1])
	m := w * h
	for x1 < x2 && x1 >= 0 {
		w = x2 - x1
		h = min(height[x2], height[x1])
		if w*h > m {
			m = w * h
		}
		if height[x1] < height[x2] {
			x1++
		} else {
			x2--
		}
	}
	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
