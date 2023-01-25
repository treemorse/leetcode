package main

func rotate(nums []int, k int) {
	result := []int{}
	n := len(nums) - (k % len(nums))

	result = append(nums[n:len(nums)], nums[:n]...)

	copy(nums, result)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	rotate(nums, 3)
}
