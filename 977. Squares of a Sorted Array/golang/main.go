package main

import "fmt"

func sortedSquares(nums []int) []int {
	var ans []int

	for _, num := range nums {
		ans = append(ans, num*num)
	}

	for i := range ans {
		for j := 0; j < len(ans)-i-1; j++ {
			if ans[j] > ans[j+1] {
				ans[j], ans[j+1] = ans[j+1], ans[j]
			}
		}
	}

	return ans
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
