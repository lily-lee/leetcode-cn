package _1137

func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	total, a, b, c := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		total = a + b + c
		a = b
		b = c
		c = total
	}
	return total
}
