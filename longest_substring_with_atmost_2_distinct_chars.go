package main

import "fmt"

func incCharCount(cnt *map[byte]int, ch byte) (isDuplCh bool) {
	chCnt, isDuplCh := (*cnt)[ch]
		isDuplCh = !(!isDuplCh || (chCnt == 0))
	chCnt++
	(*cnt)[ch] = chCnt
	return isDuplCh
}

func unsetCharCount(cnt *map[byte]int, ch byte) (isReducedUniqChar bool) {
	chCnt, _ := (*cnt)[ch]
	chCnt--
	isReducedUniqChar = (chCnt == 0)
	(*cnt)[ch] = chCnt
	return isReducedUniqChar
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	leftIdx, rightIdx := 0, 0
	currLeftIdx, currRightIdx := leftIdx, rightIdx

	distinctCount := 0
	charCounts := make(map[byte]int)

	for (currLeftIdx <= currRightIdx) && (currRightIdx < len(s)) {
		isDuplCh := incCharCount(&charCounts, s[currRightIdx])
		if !isDuplCh {
			distinctCount++
		}

		if distinctCount > 2 {
			// unset char counts and modify distinctCount
			isReducedUniqChar := unsetCharCount(&charCounts, s[currLeftIdx])
			if isReducedUniqChar {
				distinctCount--
			}
			isReducedUniqChar = unsetCharCount(&charCounts, s[currRightIdx])
			if isReducedUniqChar {
				distinctCount--
			}

			currLeftIdx += 1
		} else if distinctCount == 2 {
			if (rightIdx - leftIdx) < (currRightIdx - currLeftIdx) {
				leftIdx = currLeftIdx
				rightIdx = currRightIdx
			}
			currRightIdx++
		} else {
			currRightIdx++
		}
	}

	if (rightIdx == 0) {
		return len(s)
	}

	return rightIdx - leftIdx + 1
}

func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("bbbbbbbbbbb"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("bbbbbbbbbbbbbbbba"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("abaaaaa"))
	fmt.Println(lengthOfLongestSubstringTwoDistinct(""))
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb"))
}