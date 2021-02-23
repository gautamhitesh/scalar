// make the pattern
//        1
//      2 1
//	  3 2 1
//	4 3 2 1

// Explanation: print one row for a value of i
// 4 3 2 1, here i=4. So, for j=i+1,j>0;j-- print j

package main

import "fmt"

func main() {
	row := 5
	arr := make([][]int, row)

	for i := 0; i < row; i++ {
		arr[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := row - 1; j >= 0; j-- {
			if j >= row-i-1 {
				arr[i][j] = row - j
			} else {
				arr[i][j] = 0
			}
		}
	}
	fmt.Println(arr)
}

// func printArray(A [][4]int) {
// 	for i := 0; i < row; i++ {
// 		for j := 0; j > col; j++ {
// 			fmt.Print(A[i][j])
// 		}
// 		fmt.Print("\n")
// 	}
// }
