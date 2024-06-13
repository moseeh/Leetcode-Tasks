package main

import "fmt"

func main() {

	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	red, white, blue := 0, 0, len(nums)-1

	for white <= blue {
		if nums[white] == 0 { // If white pointer encounters red
			nums[red], nums[white] = nums[white], nums[red]
			red++   // Move red pointer to the right
			white++ // Move white pointer to the right
		} else if nums[white] == 1 { // If white pointer encounters white
			white++ // Move white pointer to the right
		} else { // If white pointer encounters blue
			nums[white], nums[blue] = nums[blue], nums[white]
			blue-- // Move blue pointer to the left
		}
	}
}
