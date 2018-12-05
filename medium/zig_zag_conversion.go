package mediums

// https://leetcode-cn.com/problems/zigzag-conversion/description/
func Convert(s string, numRows int) string {
	// 思路：维护一个游标，用来记录是向上还是向下存储，利用 slice 的长度可变性
	var down, up bool
	res := [][]string{}
	for i := 0; i < numRows; i++ {
		res = append(res, []string{})
	}
	cursor := 0
	for i := 0; i < len(s); i++ {
		if cursor == 0 {
			down = true
			up = false
		}
		if cursor == numRows-1 {
			up = true
			down = false
		}
		res[cursor] = append(res[cursor], string(s[i]))
		if down {
			cursor++
		}
		if up {
			cursor--
			if cursor < 0 {
				cursor = 0
			}
		}
	}
	str := []string{}
	for i := 0; i < len(res); i++ {
		str = append(str, res[i]...)
	}
	ret := ""
	for _, v := range str {
		ret += v
	}
	return ret
}
