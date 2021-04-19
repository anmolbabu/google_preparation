package main

import "fmt"

type QueueElement struct {
	Value int
	Index int
}

type Queue []QueueElement

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Insert(ele int, idx int) {
	(*q) = append(
		*q,
		QueueElement{
			Value: ele,
			Index: idx,
		},
	)
}

func (q *Queue) DeQueue() QueueElement {
	retVal := (*q)[0]
	(*q) = (*q)[1:]
	return retVal
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
/*
func canJump(nums []int) bool {
	return bfs(nums, 0)
}
*/
func bfs(nums []int, currIdx int) bool {
	q := NewQueue()
	visited := make(map[int]interface{})

	q.Insert(nums[currIdx], currIdx)
	visited[currIdx] = nil

	for !q.IsEmpty() {
		currQueueEle := q.DeQueue()

		if currQueueEle.Index == len(nums) - 1 {
			return true
		}

		for _, neighbour := range getNeighbours(nums, currQueueEle.Index) {
			if _, neighbourVisited := visited[neighbour]; !neighbourVisited {
				if !((neighbour < (len(nums) - 1)) && (nums[neighbour] == 0)) {
					q.Insert(nums[neighbour], neighbour)
					visited[neighbour] = nil
				}
			}
		}
	}

	return false
}

func getNeighbours(nums []int, currIdx int) []int {
	jumpIdxs := []int{}

	maxJumpsFromCurrIdx := nums[currIdx]
	for idx := 1; (idx <= maxJumpsFromCurrIdx) && ((currIdx + idx) < len(nums)); idx++ {
		jumpIdxs = append(jumpIdxs, currIdx + idx)
	}

	return jumpIdxs
}

func canJump(nums []int) bool {
	leftMostGoodIdx := len(nums) - 1

	for idx := len(nums) - 2; idx >= 0; idx-- {
		if idx + nums[idx] >= leftMostGoodIdx {
			leftMostGoodIdx = idx
		}
	}

	return (leftMostGoodIdx == 0)
}

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
	fmt.Println(canJump([]int{3,2,1,0,4}))
	fmt.Println(canJump([]int{2,0}))
}