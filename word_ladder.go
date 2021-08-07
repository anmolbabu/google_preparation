/*
A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

Every adjacent pair of words differs by a single letter.
Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
sk == endWord
Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.



Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.
Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.


Constraints:

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord, endWord, and wordList[i] consist of lowercase English letters.
beginWord != endWord
All the words in wordList are unique.
 */

package main

import "fmt"

type WordNode struct {
	Word string
	Level int
}

type WordNodeBFSQueue []WordNode

func NewBFSQueue() WordNodeBFSQueue {
	return []WordNode{}
}

func (queue *WordNodeBFSQueue) Insert(word string, level int) {
	*queue = append(*queue, WordNode{word, level})
}

func (queue *WordNodeBFSQueue) DeQueue() *WordNode {
	var ele *WordNode

	if len(*queue) > 0 {
		ele = &(*queue)[0]
	}

	newStartIdx := 0
	if len(*queue) > 0 {
		newStartIdx = 1
	}

	*queue = (*queue)[newStartIdx:]

	return ele
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !isWordExists(endWord, wordList) {
		return 0
	}

	visistedWordToLevelInStart := make(map[string]int)
	visitedWordToLevelInEnd := make(map[string]int)

	startBFSQueue := NewBFSQueue()
	startBFSQueue.Insert(beginWord, 1)
	visistedWordToLevelInStart[beginWord] = 1

	endBFSQueue := NewBFSQueue()
	endBFSQueue.Insert(endWord, 1)
	visitedWordToLevelInEnd[endWord] = 1

	allWordNeighbours := make(map[string][]string)

	for _, currWord := range wordList {
		for wordChIdx := 0; wordChIdx < len(currWord); wordChIdx++ {
			newWord := fmt.Sprintf("%s*%s", currWord[:wordChIdx], currWord[wordChIdx+1:])

			if newWordMappings, ok := allWordNeighbours[newWord]; ok {
				allWordNeighbours[newWord] = append(newWordMappings, currWord)
			} else {
				allWordNeighbours[newWord] = []string{currWord}
			}
		}
	}

	var currLevel int
	for (len(startBFSQueue) > 0) && (len(endBFSQueue) > 0) {
		currLevel += 1

		startNode := startBFSQueue.DeQueue()
		foundHeight := visitNodes(startNode, allWordNeighbours, &startBFSQueue, visistedWordToLevelInStart, visitedWordToLevelInEnd)
		if foundHeight != -1 {
			return foundHeight
		}

		endNode := endBFSQueue.DeQueue()
		foundHeight = visitNodes(endNode, allWordNeighbours, &endBFSQueue, visitedWordToLevelInEnd, visistedWordToLevelInStart)
		if foundHeight != -1 {
			return foundHeight
		}
	}

	return 0
}

func visitNodes(currWord *WordNode, allWordNeighbours map[string][]string, queue *WordNodeBFSQueue, visitedWords, otherVisitedWords map[string]int) int {
	for idx := 0; idx < len(currWord.Word); idx++ {
		newWord := fmt.Sprintf("%s*%s", currWord.Word[:idx], currWord.Word[idx+1:])

		neighbours, ok := allWordNeighbours[newWord]

		if !ok || len(neighbours) == 0 {
			continue
		}

		for _, currAdjWord := range neighbours {
			if prevLevel, ok := otherVisitedWords[currAdjWord]; ok {
				return prevLevel + currWord.Level
			}

			if _, ok := visitedWords[currAdjWord]; !ok {
				queue.Insert(currAdjWord, currWord.Level + 1)
				visitedWords[currAdjWord] = currWord.Level + 1
			}
		}
	}

	return -1
}

func isWordExists(search string, wordList []string) bool {
	for _, currWord := range wordList {
		if search == currWord {
			return true
		}
	}

	return false
}