package solution130

import "fmt"

func solve(board [][]byte) {
	row := len(board)
	col := len(board[0])

	if row < 2 || col < 2 {
		return
	}

	keepMatrix := make([][]bool, row)
	for i := range keepMatrix {
		keepMatrix[i] = make([]bool, col)
	}

	for i := 0; i < row; i++ {
		if board[i][0] == 'O' {
			updateKeepMatrix(i, 0, row, col, keepMatrix, board)
		}
		if board[i][col-1] == 'O' {
			updateKeepMatrix(i, col-1, row, col, keepMatrix, board)
		}
	}
	for i := 1; i < col-1; i++ {
		if board[0][i] == 'O' {
			updateKeepMatrix(0, i, row, col, keepMatrix, board)
		}
		if board[row-1][i] == 'O' {
			updateKeepMatrix(row-1, i, row, col, keepMatrix, board)
		}
	}

	fmt.Println(keepMatrix)

	// 再次遍历以合并KeepMatrix和board两个矩阵得到新矩阵
	for i := range board {
		for j := range board[i] {
			if !keepMatrix[i][j] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// 递归更新keepMartix
func updateKeepMatrix(row, col, rowCount, colCount int, keepMatrix [][]bool, board [][]byte) {
	if row < 0 || row > rowCount-1 || col < 0 || col > colCount-1 {
		return
	}

	if getKeepMatrix(row, col, rowCount, colCount, keepMatrix) {
		return
	}

	atEdge := row == 0 || col == 0 || row == rowCount-1 || col == colCount-1

	if board[row][col] == 'O' {
		if atEdge {
			keepMatrix[row][col] = true
		} else {
			keepMatrix[row][col] =
				getKeepMatrix(row-1, col, rowCount, colCount, keepMatrix) ||
					getKeepMatrix(row+1, col, rowCount, colCount, keepMatrix) ||
					getKeepMatrix(row, col-1, rowCount, colCount, keepMatrix) ||
					getKeepMatrix(row, col+1, rowCount, colCount, keepMatrix)
		}
	}
	if keepMatrix[row][col] {
		updateKeepMatrix(row-1, col, rowCount, colCount, keepMatrix, board)
		updateKeepMatrix(row+1, col, rowCount, colCount, keepMatrix, board)
		updateKeepMatrix(row, col-1, rowCount, colCount, keepMatrix, board)
		updateKeepMatrix(row, col+1, rowCount, colCount, keepMatrix, board)
	}
}

func getKeepMatrix(row, col, rowCount, colCount int, keepMatrix [][]bool) bool {
	if row < 0 || row > rowCount-1 || col < 0 || col > colCount-1 {
		return false
	}

	return keepMatrix[row][col]
}
