package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("enter a valid string")
		return
	}
	s := strings.Join(arguments, " ")
	s1 := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			s1 += string('z' - (v - 'a'))
		} else if v >= 'A' && v <= 'Z' {
			s1 += string('Z' - (v - 'A'))
		} else {
			s1 += string(v)

		}
	}
	fmt.Println(s1)

}
