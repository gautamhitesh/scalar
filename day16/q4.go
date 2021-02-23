//return 1 if the array elements can be sorted in an AP

package main

import (
	"fmt"
	"sort"
)

func isAP(A []int) int {
	sort.Ints(A)
	diff := A[1] - A[0]
	for i := 1; i < len(A)-1; i++ {
		if A[i+1]-A[i] != diff {
			return 0
		}
	}
	return 1
}

func main() {
	A := []int{2, 4, 1}
	fmt.Println(isAP(A))
}
