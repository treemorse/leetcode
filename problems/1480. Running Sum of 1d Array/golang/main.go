package main

func runningSum(nums []int) []int {
	increment := 0
	result := []int{}

	for _, val := range nums {
		result = append(result, val+increment)
		increment += val
	}

	return result
}

func main() {
	runningSum(nil)
}
