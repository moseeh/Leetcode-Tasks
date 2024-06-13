package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
}

func uniqueOccurrences(arr []int) bool {
	mymap := make(map[int]int)

	for _, v := range arr {
		mymap[v]++
	}
	uniqueCount := make(map[int]bool)
	for _, count := range mymap {
		if uniqueCount[count] {
			return false
		}
		uniqueCount[count] = true
	}
	return true
}
