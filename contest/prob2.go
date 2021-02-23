package main

import (
	"fmt"
	"sort"
)

func pos(A []int, x int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == x {
			return i
		}
	}
	return 0
}
func main() {
	// arr1 := []int{412, 57, 662, 601, 405}
	// arr2 := []int{732, 199, 766, 835, 914}
	arr1 := []int{2, 4, 1}
	arr2 := []int{3, 5, 6}
	arr3 := arr1
	sort.Ints(arr3)
	arr4 := [3]int{}
	for i := 0; i < len(arr1); i++ {
		arr4[i] = arr2[pos(arr3, arr3[i])]
	}
	fmt.Println(isFeasible(arr1, arr2))
}

func isFeasible(A []int, B []int) string {

	tv1 := 0
	tv2 := 0
	for i := 0; i < len(A); i++ {
		if tv1 == 0 {
			tv1 = B[i]
		} else if tv2 == 0 {
			tv2 = B[i]
		} else {
			if tv1 < A[i] {
				tv1 = B[i]
			} else {
				if tv2 < A[i] {
					tv2 = B[i]
				} else {
					return "No"
				}
			}
		}
	}
	return "Yes"
}
