package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	var newHead *Node

	if head == nil {
		return head
	}

	currPtr := head
	for currPtr != nil {
		newNode := &Node{Val: currPtr.Val, Next: currPtr.Next}

		if newHead == nil {
			newHead = newNode
		}

		currPtr.Next = newNode

		currPtr = newNode.Next
	}

	currPtr = head
	for currPtr != nil {
		if currPtr.Random != nil {
			currPtr.Next.Random = currPtr.Random.Next
		}

		currPtr = currPtr.Next.Next
	}


	oldCurrPtr := head
	newCurrPtr := head.Next

	for oldCurrPtr != nil {
		if oldCurrPtr.Next != nil {
			oldCurrPtr.Next = oldCurrPtr.Next.Next
		}

		if newCurrPtr.Next != nil {
			newCurrPtr.Next = newCurrPtr.Next.Next
		}


		oldCurrPtr = oldCurrPtr.Next
		newCurrPtr = newCurrPtr.Next
	}

	return newHead
}

