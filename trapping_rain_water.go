package main

func trap(height []int) int {
	result := 0

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	for idx := 0; idx < len(height); idx++ {
		if idx == 0 {
			leftMax[idx] = height[idx]
		} else {
			leftMax[idx] = max(leftMax[idx - 1], height[idx])
		}
	}

	for idx := len(height) - 1; idx >= 0 ; idx-- {
		if idx == len(height) - 1 {
			rightMax[idx] = height[idx]
		} else {
			rightMax[idx] = max(rightMax[idx + 1], height[idx])
		}
	}

	for idx := 0; idx < len(height); idx++ {
		result += min(rightMax[idx], leftMax[idx]) - height[idx]
	}

	return result
}
