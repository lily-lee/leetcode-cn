package _0079

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	marked := make([][]bool, h)
	for i := range marked {
		marked[i] = make([]bool, w)
	}
	for m := 0; m < h; m++ {
		for n := 0; n < w; n++ {
			if search(board, m, n, word, 0, marked) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, m, n int, word string, index int, marked [][]bool) bool {
	if index == len(word) {
		return true
	}

	if m >= len(board) || m < 0 || n >= len(board[0]) || n < 0 {
		return false
	}

	if board[m][n] == word[index] && !marked[m][n] {
		marked[m][n] = true
		up := search(board, m-1, n, word, index+1, marked)
		down := search(board, m+1, n, word, index+1, marked)
		left := search(board, m, n-1, word, index+1, marked)
		right := search(board, m, n+1, word, index+1, marked)
		if up || down || left || right {
			return true
		} else {
			marked[m][n] = false
			return false
		}
	}
	return false
}
