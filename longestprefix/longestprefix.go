package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"hello", "hellwowow", "helloworld"}))
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
