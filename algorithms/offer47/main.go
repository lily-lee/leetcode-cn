package offer47

func maxValue(grid [][]int) int {
	for i := 0; i < len(grid)-1; i++ {
		grid[i+1][0] += grid[i][0]
	}
	for i := 0; i < len(grid[0])-1; i++ {
		grid[0][i+1] += grid[0][i]
	}
	for m := 1; m < len(grid); m++ {
		for n := 1; n < len(grid[0]); n++ {
			grid[m][n] += max(grid[m][n-1], grid[m-1][n])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
