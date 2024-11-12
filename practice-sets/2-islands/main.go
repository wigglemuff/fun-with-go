package main

import "fmt"

func numberOfIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	islands := 0

	var dfs func(x, y int)

	dfs = func(x, y int) {
		if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] != '1' {
			return
		}

		grid[x][y] = 'X'

		dfs(x, y-1)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x+1, y)

	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				islands += 1
				dfs(i, j)
			}
		}
	}
	return islands
}

func maxAreaOfIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	maxarea := 0

	var areaOfAnIsland func(x, y int) int
	areaOfAnIsland = func(x, y int) int {
		if x < 0 || y < 0 || x >= rows || y >= cols || grid[x][y] != '1' {
			return 0
		}

		grid[x][y] = 'X'

		return 1 + areaOfAnIsland(x, y-1) + areaOfAnIsland(x, y+1) + areaOfAnIsland(x-1, y) + areaOfAnIsland(x+1, y)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				area := areaOfAnIsland(i, j)
				maxarea = max(maxarea, area)
			}
		}
	}

	return maxarea
}

func main() {
	// expected: 1
	fmt.Println("Number of islands: ", numberOfIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	// expected: 3
	fmt.Println("Number of islands: ", numberOfIslands([][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '0'},
	}))
	// expected: 6
	fmt.Println("Max area of islands: ", maxAreaOfIslands([][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '0'},
	}))
}
