package mediums

var (
	strIntMap = map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/description/
func MyAtoi(str string) int {
	// 思路：根据题意，string 只能以数字、+、-开头；
	// 若是 - 开头则可以置 base 为 -1，然后仅计算正值即可
	// 在剔除开头符号（或根本不需要剔除）后，找到 string 开头的那一段的数字的起始位置和终止位置 left、right 即可计算正值
	str = trim(str, " ")
	if str == "" {
		return 0
	}
	if _, ok := strIntMap[string(str[0])]; !ok && string(str[0]) != "-" && string(str[0]) != "+" {
		return 0
	}

	base := 1

	if string(str[0]) == "-" {
		if len(str) == 1 {
			return 0
		}
		str = str[1:]
		base = -1
	} else if string(str[0]) == "+" {
		if len(str) == 1 {
			return 0
		}
		str = str[1:]
	}

	left, right := getIndexFromString(str)
	ret := 0
	for i := left; i < right; i++ {
		ret = ret*10 + strIntMap[string(str[i])]
		if ret > 2<<30-1 {
			if base == 1 {
				return (2<<30 - 1) * base
			}
			return 2 << 30 * base
		}
	}
	return ret * base
}

func getIndexFromString(str string) (int, int) {
	for i, v := range str {
		if _, ok := strIntMap[string(v)]; !ok {
			return 0, i
		}
	}
	return 0, len(str)
}

func trim(str string, target string) string {
	i, j := 0, len(str)-1
	for i <= j {
		ok := false
		if string(str[i]) == target {
			i++
			ok = true
		}
		if string(str[j]) == target {
			j--
			ok = true
		}
		if !ok {
			break
		}
	}
	if i > j {
		return ""
	}
	return str[i : j+1]
}
