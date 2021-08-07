package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var carry int

	var result *ListNode

	l1CurrPtr := l1
	l2CurrPtr := l2
	resCurrPtr := result

	for (l1CurrPtr != nil) || (l2CurrPtr != nil) {
		var addend1, addend2 int

		if l1CurrPtr != nil {
			addend1 = l1CurrPtr.Val
		}

		if l2CurrPtr != nil {
			addend2 = l2CurrPtr.Val
		}

		sum := addend1 + addend2 + carry

		carry = sum / 10
		sum %= 10

		newNode := &ListNode{Val: sum}

		if resCurrPtr == nil {
			resCurrPtr = newNode
			result = resCurrPtr
		} else {
			resCurrPtr.Next = newNode
			resCurrPtr = resCurrPtr.Next
		}

		if l1CurrPtr != nil {
			l1CurrPtr = l1CurrPtr.Next
		}

		if l2CurrPtr != nil {
			l2CurrPtr = l2CurrPtr.Next
		}
	}

	if carry > 0 {
		resCurrPtr.Next = &ListNode{Val: carry}
	}

	return result
}