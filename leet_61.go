package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	oldhead := head
	n := 1
	for {
		if head.Next == nil {
			break
		}
		n++
		pre = head
		head = head.Next
	}

	k = n - k%n

	head.Next = oldhead

	for i := 0; i <= k; i++ {
		pre = head
		head = head.Next
	}
	if pre != nil {
		pre.Next = nil
	}

	return head
}
