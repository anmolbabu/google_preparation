/*
952. Largest Component Size by Common Factor
Hard

618

69

Add to List

Share
You are given an integer array of unique positive integers nums. Consider the following graph:

There are nums.length nodes, labeled nums[0] to nums[nums.length - 1],
There is an undirected edge between nums[i] and nums[j] if nums[i] and nums[j] share a common factor greater than 1.
Return the size of the largest connected component in the graph.



Example 1:


Input: nums = [4,6,15,35]
Output: 4
Example 2:


Input: nums = [20,50,9,63]
Output: 2
Example 3:


Input: nums = [2,3,6,7,4,12,21,39]
Output: 8


Constraints:

1 <= nums.length <= 2 * 104
1 <= nums[i] <= 105
All the values of nums are unique.
 */

package main

func getPrimeFactors(num int) []int {
	var res []int

	factor := 2

	for num >= factor * factor {
		if num % factor == 0 {
			res = append(res, factor)
			num /= factor
		} else {
			factor += 1
		}
	}

	res = append(res, num)

	return res
}

func max(arr []int) int {
	var max int

	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}

type UnionFind struct {
	firstFactors []int
	size []int
}

func NewUnionFind(maxVal int) *UnionFind {
	uf := &UnionFind {
		firstFactors: make([]int, maxVal + 1),
		size: make([]int, maxVal + 1),
	}

	for idx := 0; idx < len(uf.firstFactors); idx++ {
		uf.firstFactors[idx] = idx
		uf.size[idx] = 1
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.firstFactors[x] == x {
		return x
	} else {
		return uf.Find(uf.firstFactors[x])
	}

	return -1
}

func (uf *UnionFind) Union(x int, y int) {
	ux := uf.Find(x)
	uy := uf.Find(y)

	if uf.size[ux] > uf.size[uy] {
		ux, uy = uy, ux
	}

	uf.firstFactors[ux] = uy
	uf.size[uy] += uf.size[ux]
}

func largestComponentSize(nums []int) int {
	factorMap := make(map[int]int)
	uf := NewUnionFind(max(nums))
	groupCnt := make(map[int]int)
	var maxSize int

	for _, val := range nums {
		factors := getPrimeFactors(val)

		factorMap[val] = factors[0]

		for idx := 0; idx < len(factors) - 1; idx++ {
			uf.Union(factors[idx], factors[idx+1])
		}
	}

	for _, val := range factorMap {
		groupId := uf.Find(val)
		groupCnt[groupId] += 1
		maxSize = max([]int{maxSize, groupCnt[groupId]})
	}

	return maxSize
}
