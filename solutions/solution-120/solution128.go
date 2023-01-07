package solution120

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	maxStackSize := 0
	for num := range set {
		if set[num-1] {
			continue
		}
		loop := num
		for set[loop] {
			if loop-num+1 > maxStackSize {
				maxStackSize = loop - num + 1
			}
			loop++
		}
	}

	return maxStackSize
}
