package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	ll := []*ListNode{}

	for {
		if head == nil {
			break
		}
		ll = append(ll, head)
		head = head.Next
	}

	for i := right - 1; i >= left-1; i-- {
		if i == 0 {
			ll[i].Next = nil
		} else {
			ll[i].Next = ll[i-1]
		}
	}

	if left-1 >= 0 {
		if right == len(ll) {
			ll[left-1].Next = nil
		} else {
			ll[left-1].Next = ll[right]
		}

		if left-2 >= 0 {
			ll[left-2].Next = ll[right-1]
		}
	}

	if left == 1 {
		return ll[right-1]
	}

	return ll[0]
}
