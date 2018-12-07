package simples

// https://leetcode-cn.com/problems/palindrome-number/description/
func IsPalindrome(x int) bool {
	// 思路：莽
	xCopy := x

	if xCopy < 0 {
		return false
	}

	ret := 0
	for xCopy != 0 {
		divisor, remainderm := xCopy/10, xCopy%10
		ret = ret*10 + remainderm
		if ret > 2<<30-1 {
			return false
		}
		xCopy = divisor
	}
	return ret == x
}
