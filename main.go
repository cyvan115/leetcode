package main

import (
	"fmt"
	. "leetcode/medium"
)

// 本地跑一跑程序用的，莽夫要有脑子
func main() {
	list1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	list2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list3 := AddTwoNumbers(list1, list2)

	for list3 != nil {
		fmt.Println(list3.Val)
		list3 = list3.Next
	}
}
