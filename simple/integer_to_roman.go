package simples

import (
	"math"
)

var (
	intRomanMap = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}

	sortedSpecialNumbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
)

// https://leetcode-cn.com/problems/integer-to-roman/description/
func IntToRoman(num int) string {
	// 思路：intRomanMap 存储特殊值的映射，sortedSpecialNumbers 可以用来组合成特殊值
	// 转换步骤就是把数字逐一转化成罗马数字再拼凑起来
	// 找当前数字的转换值时需从 大 往 小 匹配，所以维护了 sortedSpecialNumbers
	index := 0
	arr := []string{}
	for num != 0 {
		remainder := num % 10
		val := remainder * int(math.Pow10(index))
		str := getCurrentValString(val)
		arr = append(str, arr...)
		num /= 10
		index++
	}
	ret := ""
	for i := 0; i < len(arr); i++ {
		ret += arr[i]
	}
	return ret
}

func getCurrentValString(val int) []string {
	ret := []string{}
	for val != 0 {
		cur := getCurMatch(val)
		ret = append(ret, intRomanMap[cur])
		val -= cur
	}
	return ret
}

func getCurMatch(val int) int {
	for i := len(sortedSpecialNumbers) - 1; i >= 0; i-- {
		if sortedSpecialNumbers[i] <= val {
			return sortedSpecialNumbers[i]
		}
	}
	return 0
}
