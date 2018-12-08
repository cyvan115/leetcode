package mediums

// https://leetcode-cn.com/problems/3sum/description/
func ThreeSum(nums []int) [][]int {
	// 思路：我得再捋一捋，一直超时，要换个思路， bubble sort超了我也就忍了  quick sort也超，看来是算法慢不是排序慢
	ret := [][]int{}
	quickSort(nums, 0, len(nums)-1)
	for i, _ := range nums {
		left, right := 0, len(nums)-1
		for left < right {
			if left == i || nums[left]+nums[right] < 0-nums[i] || (left != 0 && nums[left] == nums[left-1]) {
				left++
			} else if right == i || nums[left]+nums[right] > 0-nums[i] || (right != len(nums)-1 && nums[right] == nums[right+1]) {
				right--
			} else {
				var newArr []int
				if left > i {
					newArr = []int{nums[i], nums[left], nums[right]}
				} else if right < i {
					newArr = []int{nums[left], nums[right], nums[i]}
				} else {
					newArr = []int{nums[left], nums[i], nums[right]}
				}
				if !containsArray(newArr, ret) {
					ret = append(ret, newArr)
				}
				left++
				right--
			}
		}
	}
	return ret
}

func containsArray(arr []int, exists [][]int) bool {
	for _, v := range exists {
		if arrItemEqual(arr, v) {
			return true
		}
	}
	return false
}

func arrItemEqual(arr []int, target []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != target[i] {
			return false
		}
	}
	return true
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func quickSort(arr []int, left, right int) {
	i, j := left, right

	if i < j {
		base := arr[i]
		for i < j {
			for i < j && arr[j] >= base {
				j--
			}
			for i < j && arr[i] <= base {
				i++
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		quickSort(arr, left, j-1)
		quickSort(arr, j+1, right)
	}
}
