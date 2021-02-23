//return 1 if all the characters of the string are alphabets [a-z,A-Z] else return 0

//Hint:  a=97 z=122 A=65 Z=90

package main

import "fmt"

func isAlpha(A []byte) int {
	for i := 0; i < len(A); i++ {
		if (A[i] >= 65 && A[i] <= 90) || (A[i] >= 97 && A[i] <= 122) {
			continue
			// fmt.Println(A[i], "1")
		} else {
			// fmt.Println(A[i], "0")
			return 0
		}
	}
	return 1
}

func main() {
	// A := []byte{'S', 'c', 'a', 'l', 'e', 'r', 'A', 'c', 'a', 'd', 'e', 'm', 'y'}
	A := []byte{'S', 'c', 'a', 'l', 'e', 'r', '#', '2', '0', '2', '0'}
	// A := []byte{'S', 'c', 'a', 'l', 'e', 'r', 'A', 'c', 'a', 'd', 'e', 'm', 'y'}
	fmt.Println(isAlpha(A))
}
