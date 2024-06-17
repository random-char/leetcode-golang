package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var prev, current, last *ListNode = nil, head, head

	i := 1
	for last.Next != nil {
		if i == n {
			prev = current
			current = current.Next
		} else {
			i++
		}

		last = last.Next
	}

	if prev == nil {
		return current.Next
	}

	prev.Next = current.Next
	return head
}

func removeNthFromEndWithMap(head *ListNode, n int) *ListNode {
	m := make(map[int]*ListNode)

	i := 0
	current := head
	for current != nil {
		m[i] = current
		i++
		current = current.Next
	}

	if n == len(m)-1 {
		return m[1]
	}

	var next *ListNode = nil
	if n-1 >= 0 {
		next = m[n-1]
	}
	m[n+1].Next = next

	return head
}
