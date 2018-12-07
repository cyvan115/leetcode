package mediums

var (
	romanIntMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

// https://leetcode-cn.com/problems/roman-to-integer/description/
func RomanToInt(s string) int {
	// 除开特殊值一个一个加就行了
	ret := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i != 0 {
			cur := string(s[i])
			pre := string(s[i-1])
			if cur == "V" && pre == "I" {
				ret += 4
				i--
			} else if cur == "X" && pre == "I" {
				ret += 9
				i--
			} else if cur == "L" && pre == "X" {
				ret += 40
				i--
			} else if cur == "C" && pre == "X" {
				ret += 90
				i--
			} else if cur == "D" && pre == "C" {
				ret += 400
				i--
			} else if cur == "M" && pre == "C" {
				ret += 900
				i--
			} else {
				ret += romanIntMap[cur]
			}
		} else {
			ret += romanIntMap[string(s[i])]
		}
	}
	return ret
}
