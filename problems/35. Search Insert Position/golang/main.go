package main

import "fmt"

func searchInsert(nums []int, target int) int {
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

	return low
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))
}
