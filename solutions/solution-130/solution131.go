package solution130

func partition(s string) [][]string {
	ans := [][]string{}
	for i := 1; i <= len(s); i++ {
		partitionResolve(s, 0, i, []string{}, &ans)
	}

	return ans
}

// 使用回溯先试试
func partitionResolve(s string, start, end int, strs []string, ans *[][]string) {
	subStr := s[start:end]
	if !isPlalindrome(subStr) {
		return
	}

	strs = append([]string{}, strs...)
	strs = append(strs, subStr)

	if end == len(s) {
		*ans = append(*ans, strs)
		return
	}

	for sublen := 1; end+sublen <= len(s); sublen++ {
		partitionResolve(s, end, end+sublen, strs, ans)
	}
}

// 判断是否为回文串
func isPlalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[length-1-i] != str[i] {
			return false
		}
	}
	return true
}
