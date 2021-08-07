package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fastPtr, slowPtr, prevPtr *ListNode

	slowPtr, fastPtr = head, head

	for idx := 0; idx < n - 1; idx++ {
		if fastPtr == nil {
			return head
		}

		fastPtr = fastPtr.Next
	}

	for fastPtr.Next != nil {
		prevPtr = slowPtr
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}

	if prevPtr != nil {
		prevPtr.Next = slowPtr.Next
	} else {
		head = head.Next
	}

	return head
}

