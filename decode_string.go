/*
394. Decode String
Medium

5627

255

Add to List

Share
Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].



Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"
Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"
Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"
Example 4:

Input: s = "abc3[cd]xyz"
Output: "abccdcdcdxyz"


Constraints:

1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].
 */

package main

import "fmt"

type IntStack []int

type StringStack []string

func NewIntStack() *IntStack {
	return &IntStack{}
}

func (is *IntStack) Push(cnt int) {
	*is = append(*is, cnt)
}

func (is *IntStack) Pop() int {
	if len(*is) == 0 {
		return -1
	}

	poppedEle := (*is)[len(*is) - 1]
	*is = (*is)[:len(*is) - 1]

	return poppedEle
}

func NewStringStack() *StringStack {
	return &StringStack{}
}

func (is *StringStack) Push(str string) {
	*is = append(*is, str)
}

func (is *StringStack) Pop() string {
	if len(*is) == 0 {
		return ""
	}

	poppedEle := (*is)[len(*is) - 1]
	*is = (*is)[:len(*is) - 1]

	return poppedEle
}

func isNum(d byte) bool {
	return d >= '0' && d <= '9'
}

func appendNum(num int, digit byte) int {
	digitNum := int(digit - '0')
	return (num * 10) + digitNum
}

func isChar(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func appendStrByte(s string, ch byte) string {
	return fmt.Sprintf("%s%c", s, ch)
}

func decodeString(s string) string {
	is := NewIntStack()
	ss := NewStringStack()

	var currNum int
	var currStr string

	for idx := 0; idx < len(s); idx++ {
		if isNum(s[idx]) {
			currNum = appendNum(currNum, s[idx])
		} else if s[idx] == '[' {
			is.Push(currNum)

			ss.Push(currStr)

			currStr = ""
			currNum = 0
		} else if s[idx] == ']' {
			var res string
			decodedStr := ss.Pop()
			count := is.Pop()

			for repeatIdx := 0; repeatIdx < count; repeatIdx++ {
				res = string(append([]byte(res), currStr[:]...))
			}

			currStr = fmt.Sprintf("%s%s", decodedStr, res)
		} else {
			currStr = appendStrByte(currStr, s[idx])
		}
	}

	return currStr
}

