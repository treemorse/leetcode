package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1
		}

		if guess < target {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9)) // 4
}
