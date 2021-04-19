package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var (
		minNode *ListNode
		newHead *ListNode
		prevNode *ListNode
		minNodeIdx = -1
	)

	for {
		for listIdx, listHead := range lists {
			if listHead == nil {
				continue
			}

			if minNode == nil {
				minNode = listHead
				minNodeIdx = listIdx
			} else if minNode.Val > listHead.Val {
				minNode = listHead
				minNodeIdx = listIdx
			}

		}

		if (minNodeIdx != -1) && (len(lists) >= minNodeIdx) && (lists[minNodeIdx] != nil) {
			lists[minNodeIdx] = lists[minNodeIdx].Next
		}

		if prevNode != nil {
			prevNode.Next = minNode
		}

		if minNode == nil {
			break
		}

		if newHead == nil {
			newHead = minNode
		}

		prevNode = minNode

		minNode = nil
	}

	return newHead
}

// Prioirty Queue Implementation - Pending