package main

/**
n1 := &Node138{
	Val: 7,
}
n2 := &Node138{
	Val: 13,
}
n3 := &Node138{
	Val: 11,
}
n4 := &Node138{
	Val: 10,
}
n5 := &Node138{
	Val: 1,
}

n1.Next = n2

n2.Next = n3
n2.Random = n1

n3.Next = n4
n3.Random = n5

n4.Next = n5
n4.Random = n3

n5.Random = n1

fmt.Println(copyRandomList138(n1))
*/

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList138(head *Node138) *Node138 {
	if head == nil {
		return nil
	}

	list := []*Node138{}
	mm := map[*Node138]int{}

	l1 := head
	i := 0
	for {
		if l1 == nil {
			break
		}
		list = append(list, l1)
		mm[l1] = i
		i++
		l1 = l1.Next
	}

	ll := []*Node138{}
	l2 := &Node138{
		Val: list[0].Val,
	}
	ll = append(ll, l2)
	for j := 1; j < len(list); j++ {
		t := &Node138{
			Val: list[j].Val,
		}
		ll = append(ll, t)
		l2.Next = t
		l2 = t
	}

	for j := 0; j < len(list); j++ {
		n := list[j].Random
		if n == nil {
			continue
		}
		k := mm[n]

		ll[j].Random = ll[k]
	}

	return ll[0]
}
