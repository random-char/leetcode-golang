package list

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		head = removeElements(head.Next, val)
	} else {
		head.Next = removeElements(head.Next, val)
	}

	return head
}
