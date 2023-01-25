package main

import "fmt"

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(firstBadVersion(4))
}
