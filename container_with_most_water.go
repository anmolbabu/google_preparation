package main

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxArea(height []int) int {
	maxArea := 0
	leftIdx := 0
	rightIdx := len(height) - 1

	for leftIdx < rightIdx {
		maxArea = max(maxArea, min(height[leftIdx], height[rightIdx]) * (rightIdx-leftIdx))
		if height[leftIdx] < height[rightIdx] {
			leftIdx++
		} else {
			rightIdx--
		}
	}

	return maxArea
}

