//find the number of occurences of bob in the string A

package main

import "fmt"

func numBob(A string) int {
	counter := 0
	for i := 0; i < len(A)-2; i++ {
		if A[i] == 98 {
			if A[i+1] == 111 {
				if A[i+2] == 98 {
					counter++
				}
			}
		}
		// fmt.Print(A[i])
	}
	return counter

}

func main() {
	A := "bobob"
	fmt.Println(numBob(A))
}
