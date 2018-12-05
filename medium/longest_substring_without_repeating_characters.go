package mediums

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
func LengthOfLongestSubstring(s string) int {
	// 思路：map 用来记录已经在 substring 中的 character，res 记录已当前位为终止位且满足题意的最大长度
	// right 从左向右扫描，扫描完即结束，当 right 当前位已在 map 中出现时（ok && true），需要找到 left 的新位置，即当前 substring 中与 right 位值相同的下一位
	if s == "" {
		return 0
	}
	locked := make(map[string]bool, len(s))
	res := make([]int, len(s))

	res[0] = 1
	locked[charAt(s, 0)] = true

	left, right, max := 0, 1, 1

	for right < len(s) {
		// substring !contains right
		if v, ok := locked[charAt(s, right)]; !ok || !v {
			locked[charAt(s, right)] = true
		} else {
			var i int
			for i = left; i < right; i++ {
				if charAt(s, i) != charAt(s, right) {
					locked[charAt(s, i)] = false
				} else {
					i++
					break
				}
			}
			left = i
		}

		res[right] = right - left + 1
		if max < res[right] {
			max = res[right]
		}
		right++
	}
	return max
}

func charAt(s string, i int) string {
	return string(s[i])
}
