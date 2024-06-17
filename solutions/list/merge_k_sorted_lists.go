package list

func MergeKLists(lists []*ListNode) *ListNode {
	return mergeKLists(lists)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	last := len(lists) - 1
	for last > 0 {
		i, j := 0, last
		for i < j {
			lists[i] = mergeTwoLists(lists[i], lists[j])
			i++
			j--

			if i >= j {
				last = j
			}
		}
	}

	return lists[0]
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var result *ListNode

	if list1.Val < list2.Val {
		result = list1
		result.Next = mergeTwoLists(list1.Next, list2)
	} else {
		result = list2
		result.Next = mergeTwoLists(list1, list2.Next)
	}

	return result
}
