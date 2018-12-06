package mediums

// https://leetcode-cn.com/problems/container-with-most-water/description/
func MaxArea(height []int) int {
	// 思路：两边往中间加逼，因为无论移动哪一个指针，矩形的底边长度都减少 1，所以移动的目的就是改变高的较小的一方，也就是left/right 谁的 height 小移动谁
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		curHeight := height[left]
		if height[right] < height[left] {
			curHeight = height[right]
		}
		tmp := curHeight * (right - left)
		if tmp > max {
			max = tmp
		}
		if left+1 == right {
			break
		}
		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}
	return max
}
