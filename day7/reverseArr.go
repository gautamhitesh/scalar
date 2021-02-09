package main

import "fmt"

func main() {
	fmt.Println("Reverse array")
	var arr = []int{1, 2, 3, 4, 5, 6}
	lenArr := len(arr)
	for i := 0; i < lenArr/2; i++ {
		temp := arr[i]
		arr[i] = arr[lenArr-1-i]
		arr[lenArr-1-i] = temp
	}
	fmt.Println("Result", arr)
}
