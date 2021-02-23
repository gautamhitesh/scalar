// prob1
// reverse the string e.g S:= The sky is blue
// ans: bllue is sky The

package main

import (
	"fmt"
	"strings"
)

func reverseString(A string) string {
	A = strings.TrimSpace(A) //level2
	temp := ""
	end := len(A) - 1
	spaces := 0
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == ' ' {
			spaces++
			if spaces <= 1 {
				for counter := i + 1; counter < end; counter++ {
					temp = temp + fmt.Sprintf("%c", A[counter])
				}
				temp = temp + " "
			}
			end = i

		} else if A[i] != ' ' {
			spaces = 0
		}
		if i == 0 {
			for counter := 0; counter < end; counter++ {
				temp += fmt.Sprintf("%c", A[counter])
			}
		}
	}
	return temp

}

func main() {
	// S := "The sky is blue"
	//S := "       fwbpudnbrozzifml osdt ulc jsx kxorifrhubk ouhsuhf sswz qfho dqmy sn myq igjgip iwfcqq                 "
	S := "qxkpvo  f   w vdg t wqxy ln mbqmtwwbaegx   mskgtlenfnipsl bddjk znhksoewu zwh bd fqecoskmo"
	fmt.Println(reverseString(S))
}

//level 2 is removing trailing and leading spaces
// level 3 is if there are extra spaces between words reduce them to one
