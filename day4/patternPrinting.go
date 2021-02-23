//print the traingle pattern but with array notations for each line as an array
//e.g A:=3
// [
//	[1]
// [1,2]
//	[1,2,3]
//]

//return a 2 d array
//first row will have one element, 2nd row will have two and so on

package main

import (
	"fmt"
)

func twoDArray(A int) [][]int {
	X := make([][]int, A)
	row := 0
	for i := 0; i < A; i++ {
		X[i] = make([]int, row+1)
		row++
	}
	for i := 0; i < A; i++ {
		for j := 0; j <= i; j++ {
			X[i][j] = j + 1
		}
	}

	return X
}
func main() {
	A := 4

	fmt.Println(twoDArray(A))
}
