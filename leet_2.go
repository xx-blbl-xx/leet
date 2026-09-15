package main

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	p := 0
	n := 0

	num := l1.Val + l2.Val
	p = num / 10
	n = num % 10

	l3 := &ListNode{
		Val: n,
	}
	head := l3
	l1 = l1.Next
	l2 = l2.Next
	for {
		if l1 == nil && l2 == nil {
			break
		}

		l1n := 0
		l2n := 0
		if l1 != nil {
			l1n = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2n = l2.Val
			l2 = l2.Next
		}

		num := l1n + l2n + p
		p = num / 10
		n = num % 10
		l3.Next = &ListNode{
			Val: n,
		}
		l3 = l3.Next

	}

	if p != 0 {
		l3.Next = &ListNode{Val: p}
	}
	return head
}
