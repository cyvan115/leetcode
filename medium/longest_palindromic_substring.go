package mediums

// https://leetcode-cn.com/problems/longest-palindromic-substring/description/
func LongestPalindrome(s string) string {
	// 思路：依次遍历，分奇数长度与偶数长度两种情况考虑回文串，取出两者最大值再向后遍历
	if s == "" {
		return ""
	}
	index := 0
	max := 1
	for i := 0; i < len(s); i++ {
		cnt1 := oddPalindrome(i, s)
		cnt2 := evenPalindrome(i, s)
		var cnt int
		if cnt1 > cnt2 {
			cnt = cnt1
		} else {
			cnt = cnt2
		}
		if cnt > max {
			max = cnt
			index = i
		}
	}
	if max%2 != 0 {
		return s[index-max/2 : index+max/2+1]
	}
	return s[index-max/2+1 : index+max/2+1]
}

func oddPalindrome(i int, s string) int {
	cnt := 1
	left, right := i-1, i+1
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		cnt += 2
	}
	return cnt
}

func evenPalindrome(i int, s string) int {
	cnt := 0
	left, right := i, i+1
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		cnt += 2
	}
	return cnt
}
