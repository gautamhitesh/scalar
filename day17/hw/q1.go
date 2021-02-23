// find the number of amazing sub strings of S
// amazing substring is the one that starts with vowels
// like subarrays, substrings maintain order amd continuity

//e.g ABEC: total substrings are
// A,B,E,C, AB, ABE, ABEC, BE, BEC, EC

//amazing substrings are:
// A, E, AB, EC, ABEC, ABE

// solution:
// since conitnuity is maintained 2nd counter will be next to i
// ans=Sum(vowels*characters after it)+number of vowels
// ans= (3)+(1)+1+1

//Hint: vowels ACII 65-97,70-102,74-106,80-112,86-118
//count number of vowels and characters after them
//hashing (not required)
package main

import "fmt"

func numAmazingSubstrings(A string) int {
	countVowels := 0
	charAfterVowels := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 65 || A[i] == 97 || A[i] == 69 || A[i] == 101 || A[i] == 73 || A[i] == 105 || A[i] == 79 || A[i] == 111 || A[i] == 85 || A[i] == 117 {
			countVowels++
			charAfterVowels += len(A) - 1 - i
		}
		// fmt.Println(A[i])
	}
	// fmt.Println("Vowels ", countVowels)
	// fmt.Println("characters after Vowels ", charAfterVowels)
	return (countVowels + charAfterVowels) % 10003
}

func main() {
	A := "AzZGBauYuTknYjjWEEbLvqMQwnoSgXKBdHKEVpeoxYNNtBfrxQrPauttMzjKFayKwMeuChjzCocgAHfAmPCatOqarzLuabyTnxoheeocVshfERNssVPfRyPmwvOVGbzNAuvryYNWwIeyZLMlLbkcFFJRHjEIgIwOThRJJFpLbunVrbhAYsMtdsKslLAGElwrZjvZweIytMpPEYVmktQeNleNROcTjrNxXeHvOMMTMfqZHTUjetojHFzZwOekfAILYISANxeJFRNgeZDKoOTddXqxduPIjGXsRSSkgIqKMeSqlQwAKtdrkvHLgmKleNAPEztGMPmZzUuCImLAhzUnXmsVzFvJUTeIKleuRnMPNAPWJYAZLXgzTBPmkXVShbBSlIAJSeFgvRFvcoqfVFgHUefxUxuYFCfUxbxlOslUhYRFXKmMmqDIQhhfXyGqUwGMSYeLLpEsKAhvFUzavDCOUgtmmNMnsvfmCdPWuWIjuUfZCBTrWnaDopbqXcjzSqRMpQWIBNnMcOQZjDkjPkxsuwENYQyjgSHFJrgSLnwbEInBfdeIfBbVuZZbBrblJgKHOmLZACLQkSRxxQJeUMPIQutraxFtrRcSeqAejOTSqaFGglQOoWPkcNOnLIgfclWNtjGQRVMlqCPUnUlOLbHfkzUyNAmTsswWtZjIGUBrLmRmstHgVcRUgWdQTCEPRzjPVTMJRJocYHftwoRzOSyQexjmceRHdqFdgNuGmGTUdXQaNwKmvOUzZPzGCBVcbVLgMoQrESbpVGteVVntOwEWxXsZnSAoIfBSsWVhDFBuDTkcrnsPdmmSHymouxHlcgtjgKUAPznxsIRUjDFsrjadJjEtPaWTVBHpatqYeSgrpWJDOGfgIGQPcTIXVsCVyCfKMpcXWGkvwuRuTmvCbNMLeUkZrEpYZdlKAgFELfwCCbZCYBcXhfUrsIHfdwhYyxHKAMYMERwlyRtxObDoxBhjXmynYkmkYZrkzlCuvrhW"
	fmt.Println(numAmazingSubstrings(A), "Length of string", len(A))
}
