// package main

import(
	"fmt"
)

func singleDigit(A []int) int{
	temp := 0
	for i := 0; i < len(A); i++{
		temp=A[i]^temp
		fmt.Println(temp)
	}
	return temp
}


func main(){
	var arr=[] int{1,3,5,1,4,5}
	var res=singleDigit(arr)
	fmt.Println(res)
}
