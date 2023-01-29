package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	result := []int{}
	for i := 0; i <= len(numbers); i++ {
		for j := 1; j <= len(numbers)-1; j++ {
			if j != i {
				if numbers[i]+numbers[j] == target {
					result = append(result, i+1, j+1)
					return result
				}
			}
		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
