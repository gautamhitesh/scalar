//return min cost of trees with increasing order of height
//sum of B wrt increasing order of A
// sort A, find positions of B wrt A O(n Logn)

package main

import (
	"fmt"
)

func minCost(A []int, B []int) int {
	minA := [2][3]int{}
	fmt.Print(minA)
	// j := 0
	for i := 0; i < len(A); i++ {

	}

	fmt.Print(minA)

	return 0
}

func main() {
	arr1 := []int{1, 6, 4, 2, 6, 9}
	arr2 := []int{2, 5, 7, 3, 2, 7}
	fmt.Println("Min cost", minCost(arr1, arr2))
}
