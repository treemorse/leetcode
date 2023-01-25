package main

func pivotIndex(nums []int) int {
	left := []int{}
	right := nums

	for i := range nums {
		if sum(left) == sum(right[1:]) {
			return i
		}

		left = append(left, right[0])
		right = right[1:]
	}

	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	pivotIndex(nil)
}
