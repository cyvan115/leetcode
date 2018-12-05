package simples

// https://leetcode-cn.com/problems/reverse-integer/description/
func Reverse(x int) int {
	// 思路：先看看负不负，然后从最末位开始 *10+ret，注意题目中的 2<<30 -1
	base := 1
	if x < 0 {
		base = -1
		x = -x
	}
	ret := 0

	for x != 0 {
		divisor, remainderm := x/10, x%10
		ret = ret*10 + remainderm
		if ret > 2<<30-1 {
			return 0
		}
		x = divisor
	}
	return base * ret
}
