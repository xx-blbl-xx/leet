package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	ll := []*ListNode{}
	for {
		if head == nil {
			break
		}
		ll = append(ll, head)
		head = head.Next
	}

	if len(ll) == 1 || k == 1 {
		return ll[0]
	}

	n := len(ll) / k
	for i := 0; i < n; i++ {
		rre(ll, i*k, (i+1)*k-1)
	}

	for i := 0; i < n; i++ {
		if (i+2)*k-1 < len(ll) {
			ll[i*k].Next = ll[(i+2)*k-1]
		} else {

			if (i+1)*k == len(ll) {
				ll[i*k].Next = nil
			} else {
				ll[i*k].Next = ll[(i+1)*k]
			}
		}
	}

	return ll[k-1]
}

func rre(ll []*ListNode, l, r int) {
	for i := r; i >= l+1; i-- {
		if i == 0 {
			break
		}
		ll[i].Next = ll[i-1]
	}
}
