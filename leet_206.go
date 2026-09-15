package main

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		x := head.Next
		head.Next = pre
		pre = head
		head = x
	}

	return pre
}
