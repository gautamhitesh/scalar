//find area under the curve
package main

import (
	"fmt"
)

func findArea(A []int) string {
	triArea := (A[0] + A[len(A)-1])
	trapArea := 0.0
	for i := 0; i < len(A)-1; i++ {
		trapArea += float64(A[i]) + float64(A[i+1])
		fmt.Println(A[i], A[i+1])
	}
	fmt.Println("Trapezoidal areas", trapArea)
	// trapArea /= 2
	fmt.Println("Triangles area", triArea)

	res := fmt.Sprintf("%.2f", (trapArea+float64(triArea))/2)
	return res
}
func main() {
	arr := []int{2, 1, 3}
	fmt.Println(findArea(arr))

}
