package solution120

func minimumTotal(triangle [][]int) int {
	for row, layer := range triangle {
		if row == 0 {
			continue
		}

		for i, v := range layer {
			if i == 0 {
				layer[i] = triangle[row-1][i] + v
			} else if i == len(layer)-1 {
				layer[i] = triangle[row-1][i-1] + v
			} else {
				layer[i] = min(triangle[row-1][i], triangle[row-1][i-1]) + v
			}
		}
	}

	min := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		if v < min {
			min = v
		}
	}
	return min
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
