package main

import "sort"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	m := map[int][]*ListNode{}

	for {
		if head == nil {
			break
		}
		if len(m[head.Val]) == 0 {
			m[head.Val] = []*ListNode{}
		}
		m[head.Val] = append(m[head.Val], head)
		head = head.Next
	}

	list := []*ListNode{}
	for k, l := range m {
		if len(l) > 1 {
			continue
		}
		list = append(list, m[k][0])
	}

	if len(list) == 0 {
		return nil
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Val < list[j].Val
	})

	for i := 0; i < len(list); i++ {
		if i == len(list)-1 {
			list[i].Next = nil
		} else {
			list[i].Next = list[i+1]
		}
	}

	return list[0]

}
