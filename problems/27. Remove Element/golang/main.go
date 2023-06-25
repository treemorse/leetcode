package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 2, 3}
	i := removeElement(nums, 1)

	fmt.Println(nums[:i])
}
