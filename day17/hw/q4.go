//find if the given string can be converted to a palindrome with only one replacement

//sol: find the differences from back and front if differences<=1 return true

package main

import "fmt"

func palindPoss(A string) string {
	i := 0
	count := 0
	for j := len(A) - 1; j > i; j-- {
		if A[i] != A[j] {
			count++
		}
		fmt.Println(A[i], A[j])
		i++
	}
	if count == 1 {
		fmt.Println("Count differences", count)
		return "Yes"
	}
	if count == 0 && len(A)%2 == 1 {
		return "Yes"
	}
	fmt.Println("Count differences", count)
	return "No"
}

func main() {
	A := "abba"
	A1 := "aba"

	fmt.Println(palindPoss(A))
	fmt.Println(palindPoss(A1))
}
