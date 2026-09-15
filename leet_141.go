package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle141(head *ListNode) bool {
	if head == nil {
		return false
	}

	i := head
	j := head.Next

	if j == nil {
		return false
	}

	for {
		i = i.Next
		if i == nil {
			return false
		}
		j = j.Next
		if j == nil {
			return false
		}
		j = j.Next
		if j == nil {
			return false
		}

		if j == i {
			break
		}
	}

	return true
}
