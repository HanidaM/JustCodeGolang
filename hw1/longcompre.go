package hw1

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	for i := 0; i < minLen; i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:minLen]
}
