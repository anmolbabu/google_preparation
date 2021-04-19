package main

import "fmt"

func oddEvenJumps(arr []int) int {
	cntGoodJumps := 0
	jumps := make(map[string]int)

	for idx, _ := range arr {
		startIdx := idx
		currJumps := 1
		for (startIdx != - 1) {
			startIdx = getNextJump(startIdx, arr, (currJumps % 2) == 0, &jumps)
			if startIdx == len(arr) - 1 {
				cntGoodJumps += 1
				break
			}
			currJumps++
		}
	}

	return cntGoodJumps
}

func getNextJump(idx int, arr []int, isEven bool, jumps *map[string]int) int {
	if idx == (len(arr) - 1) {
		return idx
	}

	nextJump, jumpAvailable := (*jumps)[fmt.Sprintf("%d-%v", idx, isEven)]
	if jumpAvailable {
		return nextJump
	}

	nextJump = -1

	for currIdx := idx + 1; currIdx < len(arr); currIdx++ {
		if isEven {
			if (arr[idx] >= arr[currIdx]) {
				if (nextJump == -1) || (arr[nextJump] < arr[currIdx]) {
					nextJump = currIdx
				}
			}
		} else {
			if (arr[idx] <= arr[currIdx]) {
				if (nextJump == -1) || (arr[nextJump] > arr[currIdx]) {
					nextJump = currIdx
				}
			}
		}
	}

	(*jumps)[fmt.Sprintf("%d-%v", idx, isEven)] = nextJump
	return nextJump
}

func main() {
	fmt.Println(oddEvenJumps([]int{10,13,12,14,15}))
}
