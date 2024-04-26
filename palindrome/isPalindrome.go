package main

import "fmt"

func isPalindrome(x int) bool {
	var n int
	n1 := x
	if x < 0 {
		return false
	}
	for x > 0 {
		n = n*10 + x%10
		x = x / 10
	}
	if n1 == n {
		return true
	} else {
		return false
	}
}

func main() {
	// Test cases
	fmt.Println(isPalindrome(121))  // Output: true
	fmt.Println(isPalindrome(-121)) // Output: false
	fmt.Println(isPalindrome(10))   // Output: false
}
