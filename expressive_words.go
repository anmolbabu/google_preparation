package main

import "fmt"

func checkIfExtension(idx int, permiRepCnt int, wordRLE RunLengthEncoding, extensionRLE RunLengthEncoding) bool {
	currCharCntInWord := wordRLE.Counts[idx]
	currCharCntInExt := extensionRLE.Counts[idx]

	if (currCharCntInExt == currCharCntInWord) || ((currCharCntInExt > currCharCntInWord) && (currCharCntInExt >= 3)) {
		return true
	}

	return false
}

func expressiveWords(S string, words []string) int {
	rleS := NewRLE(S)

	ans := 0

	for _, word := range words {
		rleWord := NewRLE(word)
		if len(rleS.Key) != len(rleWord.Key) {
			continue
		}

		if rleS.Key != rleWord.Key {
			continue
		}

		isMatchFound := true
		for idx := 0; idx < len(rleS.Key); idx++ {
			if (rleS.Key[idx] != rleWord.Key[idx]) || !checkIfExtension(idx, 3, rleWord, rleS) {
				isMatchFound = false
				break
			}
		}

		if isMatchFound {
			ans++
		}
	}

	return ans
}

type RunLengthEncoding struct {
	Key string
	Counts []int
}

func NewRLE(str string) RunLengthEncoding {
	rle := RunLengthEncoding{}
	prev := -1

	for idx := 0; idx < len(str); idx++ {
		if (idx == len(str) - 1) || str[idx] != str[idx+1] {
			rle.Key = fmt.Sprintf("%s%c", rle.Key, str[idx])
			rle.Counts = append(rle.Counts, idx - prev)
			prev = idx
		}
	}

	return rle
}

func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	fmt.Println(expressiveWords("", []string{"hello", "hi", "helo"}))
	fmt.Println(expressiveWords("zzzzzyyyyy", []string{"zzyy","zy","zyy"}))
}