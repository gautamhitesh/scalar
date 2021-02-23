//a number is colorful if the product of all subsequences of that number is unique
// Step 1: find all the possible subsequences
// Step2: put them into an array
// Step3: put their product into an array, find if any number is repeating

package main

import "fmt"

func isColorful(X int) int {
	arr := []int{}
	i := 0
	for X > 0 {
		arr = append(arr, X%10)
		X /= 10
		i++
	}
	prodArray := arr
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			prodArray = append(prodArray)
		}
	}
	fmt.Println(arr)
	return 1

}

func main() {
	num := 12

	fmt.Println(isColorful(num))
}
