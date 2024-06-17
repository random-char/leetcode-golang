package solutions

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := 0

	for i < j {
		area = max(area, min(height[i], height[j])*(j-i))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return area
}
