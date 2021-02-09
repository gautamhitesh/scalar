// package main

import(
	"fmt"
)
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(A []int) (int){
	minOdd:=0
	maxEven:=0
	for i :=0 ; i <len(A); i++{
		if Abs(A[i])%2==0{
			if maxEven==0 || maxEven<A[i]{
				maxEven=A[i]
			}
		}
		// fmt.Println(A[i]%2, A[i])
		if Abs(A[i])%2==1 {
			if minOdd==0 || minOdd>A[i]{
				minOdd=A[i]
			}
			// else{
			// 	fmt.Println("min odd is same", minOdd, A[i])
			// }
		}
		// fmt.Println(minOdd, maxEven, A[i], A[i]%2)
	}
	
	return maxEven-minOdd	
}

func main() {
	A:=[]int{-98, 54, -52, 15, 23, -97, 12, -64, 52, 85}
	fmt.Println(solve(A))
	
}