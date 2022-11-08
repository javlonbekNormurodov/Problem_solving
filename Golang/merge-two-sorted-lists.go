package main

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	answer := &ListNode{}
	if list1.Val < list2.Val {
		answer.Val = list1.Val
		answer.Next = mergeTwoLists(list1.Next, list2)
	} else {
		answer.Val = list2.Val
		answer.Next = mergeTwoLists(list1, list2.Next)
	}
	return answer
}
