package main
/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
 */

func dfs(grid [][]byte, rowIdx int, colIdx int, maxRows int, maxCols int) {
	if (rowIdx < 0) || (colIdx < 0) || (rowIdx >= maxRows) || (colIdx >= maxCols) || string(grid[rowIdx][colIdx]) == "0" {
		return
	}

	grid[rowIdx][colIdx] = []byte("0")[0]

	dfs(grid, rowIdx - 1, colIdx, maxRows, maxCols)
	dfs(grid, rowIdx + 1, colIdx, maxRows, maxCols)
	dfs(grid, rowIdx, colIdx - 1, maxRows, maxCols)
	dfs(grid, rowIdx, colIdx + 1, maxRows, maxCols)
}

func numIslands(grid [][]byte) int {
	maxRows := len(grid)
	maxCols := len(grid[0])
	noOfIslands := 0
	for rowIdx := 0; rowIdx < maxRows; rowIdx++ {
		for colIdx := 0; colIdx < maxCols; colIdx++ {
			if string(grid[rowIdx][colIdx]) == "1" {
				noOfIslands++
				dfs(grid, rowIdx, colIdx, maxRows, maxCols)
			}
		}
	}

	return noOfIslands
}

