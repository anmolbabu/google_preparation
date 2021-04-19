package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	strLen := len(s)
	if (strLen == 0) || (strLen == 1) {
		return strLen
	}

	leftIdx := 0
	rightIdx := leftIdx + 1
	uniqContiguousCharIdxMap := make(map[byte]int)

	uniqContiguousCharIdxMap[s[leftIdx]] = leftIdx

	for (rightIdx < len(s)) && (leftIdx <= rightIdx) {
		if prevSeenCharIdx, charAlreadySeen := uniqContiguousCharIdxMap[s[rightIdx]]; charAlreadySeen && (prevSeenCharIdx >= leftIdx) && (rightIdx > leftIdx){
			maxLen = max(maxLen, rightIdx - leftIdx)
			leftIdx = prevSeenCharIdx + 1
		}
		uniqContiguousCharIdxMap[s[rightIdx]] = rightIdx
		rightIdx += 1
	}

	if maxLen == 0 {
		return rightIdx - leftIdx
	}

	return max(maxLen, rightIdx - leftIdx)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("aab"))
}