package main

import (
	"sort"
	"fmt"
)

type VisitorIdx struct {
	Index int
	IsVisited bool
}

func threeSum(nums []int) [][]int {
	result := [][]int{}

	visitedMap := make(map[int]int)
	metaResult := make(map[string][]int)

	for leftIdx := 0; leftIdx < len(nums) - 2; leftIdx++ {
		for rightIdx := leftIdx + 1; rightIdx < len(nums); rightIdx++ {
			complement := 0-nums[leftIdx]-nums[rightIdx]
			if complementIdx, complementFound := visitedMap[complement]; complementFound && (complementIdx == leftIdx) {
				threeSumZero := []int{nums[leftIdx], nums[rightIdx], complement}
				sort.Ints(threeSumZero)
				metaResult[fmt.Sprintf("%d-%d-%d", threeSumZero[0], threeSumZero[1], threeSumZero[2])] = threeSumZero
			}
			visitedMap[nums[rightIdx]] = leftIdx
		}
	}

	for _, candidateRes := range metaResult {
		result = append(result, candidateRes)
	}

	return result
}
