package main

import "fmt"

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

func minimum(arr ...int) int {
	if len(arr) == 0 {
		return -1
	}

	minVal := arr[0]

	for _, ele := range arr {
		if ele < minVal {
			minVal = ele
		}
	}

	return minVal
}

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	cpyBuf := []byte{}

	// implement read below.
	return func(buf []byte, n int) int {
		charReadCnt := 0
		for n > 0 {
			tBuf := []byte{}
			currCharWindowCnt := 0
			if len(cpyBuf) > 0 {
				currCharWindowCnt = minimum(len(cpyBuf), n)
				tBuf = append(tBuf, cpyBuf[:currCharWindowCnt]...)
				cpyBuf = cpyBuf[currCharWindowCnt:]
			} else {
				charsReadFromFile := read4(buf)
				if charsReadFromFile == 0 {
					break
				}

				charsReadFromFile = minimum(charsReadFromFile, 4 - currCharWindowCnt, n)
				tBuf = append(tBuf, buf[:charsReadFromFile]...)
				currCharWindowCnt += charsReadFromFile
				cpyBuf = append(cpyBuf, buf[charsReadFromFile:]...)
			}

			buf = tBuf[:]
			fmt.Println(buf)
			charReadCnt += currCharWindowCnt
			n -= charReadCnt
		}

		fmt.Println(charReadCnt)
		return charReadCnt
	}
}
