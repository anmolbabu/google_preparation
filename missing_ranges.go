package main

import (
	"fmt"
	"math"
)

func findMissingRanges(nums []int, lower int, upper int) []string {
	result := []string{}

	if len(nums) == 0 {
		translateRange(&lower, &upper, &result)
		return result
	}
	currNumIter := 0
	rangeNum := lower

	naLowerRange, naHigherRange := math.MinInt32, math.MinInt32

	if nums[len(nums) - 1] < upper {
		nums = append(nums, upper + 1)
	}

	for rangeNum <= upper {
		if rangeNum == nums[currNumIter] {
			translateRange(&naLowerRange, &naHigherRange, &result)
			rangeNum++
			currNumIter++
		} else if (rangeNum > nums[currNumIter]) {
			for rangeNum > nums[currNumIter] {
				currNumIter++
			}
		} else {
			if naLowerRange == math.MinInt32 {
				naLowerRange = rangeNum
			} else {
				naHigherRange = rangeNum
			}
			if (nums[currNumIter] - 1) > rangeNum {
				rangeNum = (nums[currNumIter] - 1)
			} else {
				rangeNum++
			}
		}
	}

	translateRange(&naLowerRange, &naHigherRange, &result)
	return result
}

func translateRange(naLowerRange *int, naHigherRange *int, result *[]string) {
	missingRange := ""
	if *naLowerRange != math.MinInt32 {
		missingRange = fmt.Sprintf("%d", *naLowerRange)
	}

	if (*naHigherRange != math.MinInt32) && (*naLowerRange < *naHigherRange) {
		missingRange = fmt.Sprintf("%s->%d", missingRange, *naHigherRange)
	}

	// add mssingRange to result
	if missingRange != "" {
		*result = append(*result, missingRange)
	}

	*naLowerRange = math.MinInt32
	*naHigherRange = math.MinInt32
}

func main() {
	fmt.Println(findMissingRanges([]int{0,1,3,50,75}, 0, 99))
	fmt.Println(findMissingRanges([]int{0,1,3,50,75}, 76, 99))
	fmt.Println(findMissingRanges([]int{50,75}, 0, 60))
	fmt.Println(findMissingRanges([]int{50,75}, 0, 49))
	fmt.Println(findMissingRanges([]int{}, 1, 1))
	fmt.Println(findMissingRanges([]int{}, -3, -1))
	fmt.Println(findMissingRanges([]int{-1}, -1, -1))
	fmt.Println(findMissingRanges([]int{-1}, -2, -1))
}
