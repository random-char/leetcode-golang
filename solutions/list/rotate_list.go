package list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 1
	node := head
	for node.Next != nil {
		node = node.Next
		n++
	}

	//circle around
	node.Next = head

	prevToNewHead := node
	for i := 0; i < n-k%n; i++ {
		prevToNewHead = prevToNewHead.Next
	}

	newHead := prevToNewHead.Next
	prevToNewHead.Next = nil
	return newHead
}
