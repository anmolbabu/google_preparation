/*
210. Course Schedule II
Medium

4421

182

Add to List

Share
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]


Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
All the pairs [ai, bi] are distinct.
 */

package main

type Node struct {
	parentCnt int
	data int
	children map[int]*Node
}

type Topology struct {
	valToNode map[int]*Node
	graphRoots map[int]*Node
}

func NewNode(val int) *Node {
	return &Node {
		data: val,
		children: make(map[int]*Node),
	}
}

func NewTopology() *Topology {
	return &Topology {
		valToNode: make(map[int]*Node),
		graphRoots: make(map[int]*Node),
	}
}

func (tplg *Topology) IsGraphRootsEmpty() bool {
	return len(tplg.graphRoots) == 0
}

func (tplg *Topology) IsFullyTraversed() bool {
	return len(tplg.valToNode) == 0
}

func (tplg *Topology) Pop() *Node {
	var currParent *Node
	var currParentVal int

	for currParentVal, currParent = range tplg.graphRoots {
		for childVal, child := range currParent.children {
			child.parentCnt -= 1

			if child.parentCnt == 0 {
				tplg.graphRoots[childVal] = child
			}
		}
		break
	}

	if currParent != nil {
		delete(tplg.valToNode, currParentVal)
		delete(tplg.graphRoots, currParentVal)
	}

	return currParent
}

func (tplg *Topology) Add(parent int, child int) {
	if _, ok := tplg.graphRoots[child]; ok {
		delete(tplg.graphRoots, child)
	}

	parentNode, parentFound := tplg.valToNode[parent]
	if !parentFound {
		parentNode = NewNode(parent)
		tplg.graphRoots[parent] = parentNode
		tplg.valToNode[parent] = parentNode
	}

	if child > -1 {
		childNode, childFound := tplg.valToNode[child]
		if !childFound {
			childNode = NewNode(child)
			tplg.valToNode[child] = childNode
		}

		childNode.parentCnt += 1

		parentNode.children[child] = childNode
	}
}

func (tplg *Topology) Sort() []int {
	result := []int{}

	for !tplg.IsGraphRootsEmpty() {
		currNode := tplg.Pop()

		if currNode != nil {
			result = append(result, currNode.data)
		}
	}

	if !tplg.IsFullyTraversed() {
		return []int{}
	}

	return result
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	tplg := NewTopology()

	for idx := 0; idx < numCourses; idx++ {
		tplg.Add(idx, -1)
	}

	for _, currPrereq := range prerequisites {
		tplg.Add(currPrereq[1], currPrereq[0])
	}

	return tplg.Sort()
}

