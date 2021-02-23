//find the lexicographically largest string possible
// given a string s with two strings separated by _
// <A>_<B> you have to replace characters of A by characters in B

package main

import (
	"fmt"
	"strings"
)

func largestString(A string) string {
	//split String
	//golang pass by reference
	//sort greatest to smallest before the _ and after the _
	//compare and create the 3rd string with the both parts
	j := strings.Index(A, "_")
	j++

	fmt.Println(j)

	return "Yes"
}

func main() {

	A := "abb_c"
	fmt.Println(largestString(A))

}
