package mediums

// https://leetcode.com/problems/permutation-in-string/
func CheckInclusion(s1 string, s2 string) bool {
	// 思路：组合可以理解为 map(string-count) 相等
	// 而每两个相邻的 map 有 len(s1)-1 个元素是相同的，这样就可以得到所有的 map 了，再和 base 比较是否相等
	baseMap := make(map[string]int, 0)
	for _, v := range s1 {
		if _, ok := baseMap[string(v)]; ok {
			baseMap[string(v)] = baseMap[string(v)] + 1
		} else {
			baseMap[string(v)] = 1
		}
	}
	maps := make(map[int]map[string]int, 0)
	var oldChar string
	for i:=0;i<len(s2)-len(s1)+1; i++{
		var curMap map[string]int
		if i > 0 {
			curMap = maps[i-1]
			curMap[oldChar] = curMap[oldChar] - 1
			newChar := string(s2[i + len(s1) - 1])
			if _, ok := curMap[newChar]; ok {
				curMap[newChar] = curMap[newChar] + 1
			} else {
				curMap[newChar] = 1
			}
			if mapEqual(baseMap, curMap) {
				return true
			}
		} else {
			curMap = make(map[string]int, 0)
			for j := 0 ; j < len(s1); j++ {
				curChar := string(s2[j])
				if _, ok := curMap[curChar]; ok {
					curMap[curChar] = curMap[curChar] + 1
				} else {
					curMap[curChar] = 1
				}
			}
			if mapEqual(baseMap, curMap) {
				return true
			}
		}
		maps[i] = curMap
		oldChar = string(s2[i])
	}
	return false
}

func mapEqual(target, cur map[string]int) bool {
	for k, v := range target {
		if _, ok := cur[k]; ok && v == cur[k] {
			continue
		} else {
			return false
		}
	}
	return true
}