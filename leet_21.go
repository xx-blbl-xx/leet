package main

import "fmt"

func mergeTwoLists21(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	l1 := list1
	l2 := list2
	var nl *ListNode
	head := list1
	if l1.Val >= l2.Val {
		head = list2
		nl = list2
		l2 = l2.Next
	} else {
		nl = list1
		l1 = l1.Next
	}
	for {
		if l1 == nil {
			nl.Next = l2
			break
		}
		if l2 == nil {
			nl.Next = l1
			break
		}

		if nl.Val <= l1.Val && l1.Val <= l2.Val {
			nl.Next = l1
			l1 = l1.Next
			nl = nl.Next
		} else if nl.Val <= l2.Val && l2.Val <= l1.Val {
			nl.Next = l2
			l2 = l2.Next
			nl = nl.Next
		} else {
			fmt.Println(nl, l1, l2)
		}

	}

	return head
}
