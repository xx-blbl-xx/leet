package main

func removeNthFromEnd19(head *ListNode, n int) *ListNode {
	ll := []*ListNode{}
	for {
		if head == nil {
			break
		}
		ll = append(ll, head)
		head = head.Next
	}

	if len(ll) == 1 {
		return nil
	}
	if len(ll) == n {
		return ll[1]
	}

	if n == 1 {
		ll[len(ll)-n-1].Next = nil
		return ll[0]
	}

	ll[len(ll)-n-1].Next = ll[len(ll)-n+1]
	return ll[0]

}
