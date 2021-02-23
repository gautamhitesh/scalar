//find the longest common prefix string for the given array of strings

package main

import "fmt"

func longestPrefix(A []string) string {
	minLength := len(A[0])
	for i := 1; i < len(A); i++ {
		if minLength > len(A[i]) {
			minLength = len(A[i])
		}
	}
	// fmt.Println("minLength", minLength)
	count := 0
	for i := 0; i < minLength; i++ {
		colDiff := A[0][i]
		for j := 1; j < len(A); j++ {
			temp := colDiff == A[j][i]
			fmt.Println(A[j][i], colDiff)
			if temp == true {
				count++
			} else {
				break
			}
			colDiff = A[j][i]
		}
	}
	fmt.Println(count)
	count--
	temp := ""
	for i := 0; i <= count; i++ {
		temp = temp + fmt.Sprintf("%c", A[0][i])
	}
	return temp
}

func main() {

	// A := []string{"abcdefgh", "aefghijk", "abcefgh"}
	A := []string{"abab", "ab", "abcd"}
	fmt.Println(longestPrefix(A))
}
