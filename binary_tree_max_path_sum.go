package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxNum := math.MinInt32
	maxPathSumHelper(root, &maxNum)
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumHelper(root *TreeNode, maxNum *int) int {
	if root == nil {
		return 0
	}

	leftGain := max(maxPathSumHelper(root.Left, maxNum), 0)
	rightGain := max(maxPathSumHelper(root.Right, maxNum), 0)
	pathGain := root.Val + leftGain + rightGain
	*maxNum = max(*maxNum, pathGain)

	return root.Val + max(leftGain, rightGain)
}
