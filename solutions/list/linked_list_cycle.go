package list

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	checkMap := make(map[*ListNode]struct{})
	for head.Next != nil {
		if _, found := checkMap[head.Next]; found {
			return true
		}

		checkMap[head.Next] = struct{}{}
		head = head.Next
	}

	return false
}
