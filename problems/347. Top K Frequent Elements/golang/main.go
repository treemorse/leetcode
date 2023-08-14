package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	numsCount := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	result := []int{}

	for _, num := range nums {
		if _, ok := numsCount[num]; ok {
			numsCount[num]++
			continue
		}

		numsCount[num] = 1
	}

	for num, count := range numsCount {
		freq[count] = append(freq[count], num)
	}

	for i := len(freq) - 1; i >= 0; i-- {
		result = append(result, freq[i]...)
		if len(result) >= k {
			return result
		}
	}

	return result
}
