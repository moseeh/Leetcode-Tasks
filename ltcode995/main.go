package main

import (
	"fmt"
)

func minKBitFlips(A []int, K int) int {
	n := len(A)
	flip := make([]int, n)
	flipped := 0
	flips := 0

	for i := 0; i < n; i++ {
		if i >= K {
			flipped ^= flip[i-K]
		}
		if A[i] == flipped {
			if i+K > n {
				return -1
			}
			flips++
			flipped ^= 1
			if i+K < n {
				flip[i] = 1
			}
		}
	}
	return flips
}

func main() {
	A := []int{0, 1, 0}
	K := 1
	fmt.Println(minKBitFlips(A, K)) // Output: 2

	A = []int{1, 1, 0}
	K = 2
	fmt.Println(minKBitFlips(A, K)) // Output: -1

	A = []int{0, 0, 0, 1, 0, 1, 1, 0}
	K = 3
	fmt.Println(minKBitFlips(A, K)) // Output: 3
}
