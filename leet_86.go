package main

func partition86(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil || x >= 100 || x <= -100 {
		return head
	}

	l := []*ListNode{}
	r := []*ListNode{}

	for {
		if head == nil {
			break
		}

		if head.Val >= x {
			r = append(r, head)
		} else {
			l = append(l, head)
		}

		head = head.Next
	}

	if len(l) == 0 {
		return r[0]
	}
	if len(r) == 0 {
		return l[0]
	}

	for i := 0; i < len(l); i++ {
		if i == len(l)-1 {
			l[i].Next = nil
		} else {
			l[i].Next = l[i+1]
		}
	}

	for i := 0; i < len(r); i++ {
		if i == len(r)-1 {
			r[i].Next = nil
		} else {
			r[i].Next = r[i+1]
		}
	}
	l[len(l)-1].Next = r[0]

	return l[0]
}
