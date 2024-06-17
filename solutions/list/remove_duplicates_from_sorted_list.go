package list

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := head
	for node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
			continue
		}

		node = node.Next
	}

	return head
}
