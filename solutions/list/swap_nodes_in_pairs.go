package list

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	var node1, node2, tmp *ListNode

	node1 = head
	node2 = head.Next

	if node2 != nil {
		head = node2
	}

	for node1 != nil && node2 != nil {
		tmp = node2.Next
		node2.Next = node1
		node1.Next = tmp

		if prev != nil {
			prev.Next = node2
		}

		prev = node1
		node1 = node1.Next
		if node1 != nil {
			node2 = node1.Next
		} else {
			node2 = nil
		}
	}

	return head
}
