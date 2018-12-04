package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/add-two-numbers/description/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// head保存头结点方便返回，cur用来“生成”当前节点，addon用来记录进位
	head := &ListNode{}
	cur := head
	addon := 0
	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
		}
		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
		}

		val := val1 + val2 + addon
		addon = val / 10
		val = val % 10

		cur.Next = &ListNode{
			Val: val,
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		cur = cur.Next
	}
	if addon != 0 {
		cur.Next = &ListNode{
			Val: addon,
		}
	}
	return head.Next
}
