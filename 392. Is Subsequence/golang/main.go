package main

import "fmt"

func isSubsequence(s string, t string) bool {
	sIndex := 0

	if len(s) == 0 {
		return true
	}

	for i := range t {
		if s[sIndex] == t[i] {
			sIndex++
		}

		if sIndex == len(s) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
}
