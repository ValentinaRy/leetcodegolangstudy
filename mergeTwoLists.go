package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var newHead *ListNode
	if list1.Val <= list2.Val {
		newHead = list1
		list1 = list1.Next
	} else {
		newHead = list2
		list2 = list2.Next
	}
	current := newHead
	for list1 != nil || list2 != nil {
		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	return newHead
}
