package main

func nextPermutation(nums []int)  {
	swapMSBIdx := len(nums) - 2

	for (swapMSBIdx >= 0) && (nums[swapMSBIdx] >= nums[swapMSBIdx + 1]) {
		swapMSBIdx--
	}

	if swapMSBIdx >= 0 {
		swapLSBIdx := len(nums) - 1

		for (swapLSBIdx >= 0) && (nums[swapLSBIdx] <= nums[swapMSBIdx]) {
			swapLSBIdx--
		}

		nums[swapLSBIdx], nums[swapMSBIdx] = nums[swapMSBIdx], nums[swapLSBIdx]
	}

	reverseFrom(&nums, swapMSBIdx + 1)
}

func reverseFrom(nums *[]int, swapMSBIdx int) {
	leftIdx := swapMSBIdx
	rightIdx := len(*nums) - 1
	for (leftIdx < rightIdx) {
		(*nums)[leftIdx], (*nums)[rightIdx] = (*nums)[rightIdx], (*nums)[leftIdx]
		leftIdx++
		rightIdx--
	}
}

