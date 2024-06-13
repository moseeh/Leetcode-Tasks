package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	s := append([]int(nil), heights...)
	sort.Ints(s)

	var count int

	for i, v := range heights {
		if v != s[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
