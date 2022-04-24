package Leetcode

/********************************************************
* @author: Ihc
* @date: 2022/4/24 0024 11:54
* @version: 1.0
* @description:
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

*********************************************************/

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTowNumber 两数相加
func addTowNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		root   *ListNode = &ListNode{}
		cursor *ListNode = root
		carry  int
	)
	// 遍历到l1货l2的结尾,结尾处可能有进位操作
	for l1 != nil || l2 != nil || carry != 0 {
		var (
			l1Val  int
			l2Val  int
			sumVal int
		)

		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		sumVal = l1Val + l2Val + carry
		carry = sumVal / 10                    // 有进位
		sumNode := &ListNode{Val: sumVal % 10} // 无论是否进位,当前值都为 sumValue % 10
		cursor.Next = sumNode                  // 第一个cursor为空的头节点
		cursor = sumNode
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// 头节点为空,返回头节点的下个节点
	return root.Next
}
