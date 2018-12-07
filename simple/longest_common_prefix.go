package simples

// https://leetcode-cn.com/problems/longest-common-prefix/description/
func LongestCommonPrefix(strs []string) string {
	// 思路：莽
	if len(strs) == 0 {
		return ""
	}

	maxLength := len(strs[0])
	if maxLength == 0 {
		return ""
	}

	ret := ""
	for i := 0; i < maxLength; i++ {
		eq := true
		base := string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || base != string(strs[j][i]) {
				eq = false
				break
			}
		}
		if eq {
			ret += base
		} else {
			break
		}
	}
	return ret
}
