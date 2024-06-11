package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	// Create a frequency map for arr1
	freq := make(map[int]int)
	for _, num := range arr1 {
		freq[num]++
	}

	// New array to hold the sorted elements
	newarr := []int{}

	// Add elements from arr1 to newarr based on order in arr2
	for _, v2 := range arr2 {
		if count, exists := freq[v2]; exists {
			for i := 0; i < count; i++ {
				newarr = append(newarr, v2)
			}
			delete(freq, v2) // Remove the element from the map after processing
		}
	}

	// Remaining elements that are not in arr2
	remaining := []int{}
	for num, count := range freq {
		for i := 0; i < count; i++ {
			remaining = append(remaining, num)
		}
	}

	// Sort the remaining elements in ascending order
	sort.Ints(remaining)

	// Append the sorted remaining elements to newarr
	newarr = append(newarr, remaining...)

	return newarr
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2)) // Output: [2 2 2 1 4 3 3 9 6 7 19]

	arr1 = []int{28, 6, 22, 8, 44, 17}
	arr2 = []int{22, 28, 8, 6}
	fmt.Println(relativeSortArray(arr1, arr2)) // Output: [22 28 8 6 17 44]
}
