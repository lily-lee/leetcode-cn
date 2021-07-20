package _0014

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	common := ""
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		equal := true
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				return common
			}
			if strs[j][i] != c {
				equal = false
				return common
			}
		}
		if equal {
			common += string(c)
		}
	}
	return common
}
