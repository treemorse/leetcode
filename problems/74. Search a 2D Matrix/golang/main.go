package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i <= len(matrix)-1; i++ {
		if matrix[i][len(matrix[i])-1] >= target {
			return search(matrix[i], target)
		}
	}

	return false
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
