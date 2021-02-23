//String operations
//concatentate the string with itself
//remove all ppercase letters
//replace vowels with #

package main

import (
	"fmt"
)

func stringOp(A string) string {
	Arune := []rune(A)
	for i := 0; i < len(Arune); i++ {
		if Arune[i] >= 65 && Arune[i] <= 90 {
			Arune[i] = ' '
		}
		if Arune[i] == 97 || Arune[i] == 101 || Arune[i] == 105 || Arune[i] == 111 || Arune[i] == 117 {
			Arune[i] = '#'
		}
	}
	temp := ""
	for i := 0; i < len(Arune); i++ {
		if Arune[i] != ' ' {
			temp = temp + fmt.Sprintf("%c", Arune[i])
		}
	}
	temp = temp + temp
	return temp
}

//didnt work

// func removeChar(A []byte, c byte) []byte {
// 	count := 0
// 	for i := len(A) - 1; i >= 0; i-- {
// 		if A[i] == c {
// 			count++
// 			for j := i; j > 0; j-- {
// 				A[j] = A[j-1]
// 			}
// 		}
// 	}

// 	return A
// }

func main() {

	A := "AbcaZeoB"
	fmt.Println(stringOp(A))

}
