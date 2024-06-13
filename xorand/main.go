package main

import "fmt"

func getXORSum(arr1 []int, arr2 []int) int {
	var xorArr1, xorArr2 int
	for _, v := range arr1 {
		xorArr1 ^= v
	}
	for _, v := range arr2 {
		xorArr2 ^= v
	}
	return xorArr1 & xorArr2
}

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{6, 5}
	fmt.Println(getXORSum(arr1, arr2))
}
