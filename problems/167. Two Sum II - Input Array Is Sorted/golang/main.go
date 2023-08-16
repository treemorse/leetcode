package main

import "fmt"

// bad solution
//func twoSum(numbers []int, target int) []int {
//	result := []int{}
//	for i := 0; i <= len(numbers); i++ {
//		for j := 1; j <= len(numbers)-1; j++ {
//			if j != i {
//				if numbers[i]+numbers[j] == target {
//					result = append(result, i+1, j+1)
//					return result
//				}
//			}
//		}
//	}
//	return result
//}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for {
		sum := numbers[l] + numbers[r]

		if sum > target {
			r--
			continue
		}

		if sum < target {
			l++
			continue
		}

		if sum == target {
			return []int{l + 1, r + 1}
		}
	}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
