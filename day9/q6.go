//return row numbwe with max number of 1s

//count number of 1s in each row which will be N^2
//sum of rows == n2
// condition 1: each row is sorted

// traverse from top corner and increment the row as you find the 0

package main

import "fmt"

func maxRow(A [][]int) int {
	row := 0
	maxRow := 0
	col := len(A[0])
	for i:=0;i<;i++{

	
	for A[row][col] == 1 {
		col--
		maxRow = row
	}
}
	return maxRow + 1
}

func main() {
	// A := [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}

	A := [][]int{{0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 1, 1, 1, 1, 1, 1}}
	// A := [][]int{{0, 1, 1}, {0, 0, 1}, {0, 1, 1}}
	fmt.Println(maxRow(A))
}
