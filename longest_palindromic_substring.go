/*
5. Longest Palindromic Substring
Medium

13213

793

Add to List

Share
Given a string s, return the longest palindromic substring in s.



Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
Example 3:

Input: s = "a"
Output: "a"
Example 4:

Input: s = "ac"
Output: "a"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
 */
func longestPalindrome(s string) string {
results := make([][]bool, len(s))
resLeftIdx := 0
resRightIdx := resLeftIdx
prevRightIdx := 0

for idx, _ := range results {
results[idx] = make([]bool, len(s))
results[idx][idx] = true

if idx + 1 < len(s) {
results[idx][idx + 1] = (s[idx] == s[idx+1])
if results[idx][idx + 1] {
if resRightIdx - resLeftIdx  < idx + 1 + 1 - idx {
resRightIdx = idx + 1
resLeftIdx = idx
}
}
}
}

leftIdx, rightIdx := 1, 1

for rightIdx <= len(s) && leftIdx <= len(s) {
if (leftIdx == 1) && (rightIdx == len(s)) {
return s[resLeftIdx:resRightIdx+1]
}

if leftIdx == len(s) || rightIdx == len(s) {
rightIdx = prevRightIdx + 1
prevRightIdx = rightIdx
leftIdx = 0
}

if (leftIdx != rightIdx) && (leftIdx != rightIdx - 1) {
results[leftIdx][rightIdx] = (s[leftIdx] == s[rightIdx]) && results[leftIdx + 1][rightIdx - 1]
if results[leftIdx][rightIdx] && (rightIdx - leftIdx + 1) > (resRightIdx - resLeftIdx + 1) {
resRightIdx = rightIdx
resLeftIdx = leftIdx
}
}

leftIdx++
rightIdx++
}


return s[resLeftIdx: resRightIdx+1]
}