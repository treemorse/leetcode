package main

func maxArea(height []int) int {
	l, r, max, area := 0, len(height)-1, 0, 0

	for l != r {
		if height[l] < height[r] {
			area = height[l] * (r - l)
			l++
		} else {
			area = height[r] * (r - l)
			r--
		}

		if area > max {
			max = area
		}
	}

	return max
}
