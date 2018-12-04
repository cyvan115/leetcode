package simples

// https://leetcode-cn.com/problems/two-sum/description/
func TwoSum(nums []int, target int) []int {
	// 思路：怼
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
