package golang

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	count := 1

	for i := range nums {
		answer[i] = count
		count *= nums[i]
	}

	count = 1

	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= count
		count *= nums[i]
	}

	return answer
}
