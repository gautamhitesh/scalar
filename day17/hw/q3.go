//given a string A and a number B
//you can change at max B characters of A such that it minimizes
//the number of distinct characters in A
//find the number of distinct characters in A after the operation

//e.g A = "abcabbccd" B = 3

// a= 2, b=3, c=3, d=1

//replace the max number of characters so start by their lowest frequency
// ans= sum of unique elements - lowest frequencies till B
// 4 - 2 = 2
package main

import (
	"fmt"
)

func minDistinct(A string) int {
	
}

func main() {
	A := "abcabbccd"
	fmt.Println(minDistinct(A))

}
