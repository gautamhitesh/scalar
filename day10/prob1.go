// String problem
// String is a character array

// find the number of subsequences of "ag"
// abcgag : AbcGag, AbcgaG, abcgAG
// explanantion : for every A encountered in the string the answer increases by G times, where
// G are the number of Gs after A
package main

import "fmt"

func main() {
	// string1 := "abcgag"
	string1 := "abgagag"
	ans := 0
	countA := 0
	fmt.Println("Given string ", string1)
	for i := 0; i < len(string1); i++ {
		if string1[i] == 'a' {
			countA += 1
		} else if string1[i] == 'g' {
			ans += countA
		}
		// if string1[i] == 'g' {
		// 	countG++
		// } else if string1[i] == 'a' {
		// 	ans += countG
		// }
	}
	fmt.Println(ans)
}
