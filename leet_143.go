package main

func reorderList(head *ListNode) {
	arr := make([]*ListNode, 0)

	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	ll := len(arr)
	i := 0
	for ; i < ll/2; i++ {
		t := arr[i].Next
		arr[i].Next = arr[ll-i-1]
		arr[ll-i-1].Next = t
	}

	arr[i].Next = nil
}
