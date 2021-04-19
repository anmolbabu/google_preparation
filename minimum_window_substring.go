package main

import (
	"fmt"
	"math"
)

func getCharCounts(str string) (map[byte]int, int) {
	result := make(map[byte]int)
	totalCount := 0

	for idx := 0; idx < len(str); idx++ {
		result[str[idx]] += 1
		totalCount += 1
	}

	return result, totalCount
}

func getCharCount(idx int, s string, charCounts map[byte]int) int {
	if idx >= len(s) {
		return -1
	}

	cnt, ok := charCounts[s[idx]]
	if !ok {
		return -1
	}

	return cnt
}

func countCurrChar(tCharCounts map[byte]int, idx int, str string, windowCharMap *map[byte]int, windowCurrCharCount *int) {
	if idx > len(str) {
		return
	}

	reqIdxCharCnt := getCharCount(idx, str, tCharCounts)
	// > -1 means current char is in pattern
	if reqIdxCharCnt > -1 {
		(*windowCharMap)[str[idx]] += 1

		currIdxCharCnt := getCharCount(idx, str, *windowCharMap)

		if currIdxCharCnt <= reqIdxCharCnt {
			*windowCurrCharCount = *windowCurrCharCount + 1
		}
	}
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tCharCounts, charsRequired := getCharCounts(t)

	leftIdx, rightIdx := 0, 0
	minWindowLeftIdx, minWindowRightIdx := 0, math.MaxInt32

	windowCharMap := make(map[byte]int)
	windowCurrCharCount := 0
	isWindowMatch := false

	for (rightIdx < len(s)) && (leftIdx <= rightIdx) {
		for (getCharCount(leftIdx, s, tCharCounts) == -1) && (leftIdx < len(s)) {
			leftIdx++
		}

		if leftIdx > rightIdx {
			rightIdx = leftIdx
		}

		if leftIdx >= len(s) {
			break
		}

		for (windowCurrCharCount < charsRequired) && (rightIdx < len(s)) {
			countCurrChar(tCharCounts, rightIdx, s, &windowCharMap, &windowCurrCharCount)

			if windowCurrCharCount == charsRequired {
				isWindowMatch = true
				break
			}

			rightIdx++
		}

		//Check and reset window start and end indexes if relevant
		if isWindowMatch && ((rightIdx - leftIdx) < (minWindowRightIdx - minWindowLeftIdx)) {
			minWindowLeftIdx = leftIdx
			minWindowRightIdx = rightIdx
		}

		// Unset chars at:
		// 1. leftIdx: Because leftIdx changes old leftIdx char should no longer be considered
		// 2. rightIdx: Since, rightIdx is not varying, avoid rightIdx char being counted twice
		unsetChar(s, &windowCharMap, leftIdx, tCharCounts, &windowCurrCharCount)
		unsetChar(s, &windowCharMap, rightIdx, tCharCounts, &windowCurrCharCount)

		leftIdx++
		isWindowMatch = false
	}

	if minWindowRightIdx < len(s) {
		return string(s[minWindowLeftIdx : minWindowRightIdx+1])
	}

	return ""
}

func unsetChar(s string, windowCharMap *map[byte]int, idx int, tCharCounts map[byte]int, windowCurrCharCount *int) {
	if idx < len(s) {
		(*windowCharMap)[s[idx]] -= 1
		if getCharCount(idx, s, *windowCharMap) < getCharCount(idx, s, tCharCounts) {
			*windowCurrCharCount -= 1
		}
	}
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("zazbzczdzzezzzfzgzhzjzzkz", "zzz"))
	fmt.Println(minWindow("bbbbbbbba", "aa"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("a", "b"))
	fmt.Println(minWindow("ab", "b"))
	fmt.Println(minWindow("ab", "a"))
	fmt.Println(minWindow("abc", "ac"))
	fmt.Println("done")
}
