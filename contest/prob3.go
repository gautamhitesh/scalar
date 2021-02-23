//return 1 if all the elements in the array are consequetive else return 0

package main

import "fmt"

func solve(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	min := A[0]
	max := A[0]
	for j := 1; j < len(A); j++ {
		if min > A[j] {
			min = A[j]
		}
		if max < A[j] {
			max = A[j]
		}
	}
	sumFormula := len(A) * (min + max) / 2
	fmt.Println("Sum", sum)
	fmt.Println("sumFormula", sumFormula)
	if sum == sumFormula {
		fmt.Println("All consequetive")
		return 1
	} else {
		fmt.Println(sum - sumFormula)
		return 0
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(solve(arr))
}
