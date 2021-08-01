/* Given n non-negative integers a1, a2, ..., an , where each represents
a point at coordinate (i, ai). n vertical lines are drawn such that the two
endpoints of the line i is at (i, ai) and (i, 0). Find two lines, which,
together with the x-axis forms a container, such that the container contains
the most water.
*/
package main

func maxArea(height []int) int {
	if len(height) < 2 && height == nil {
		return 0
	}

	left_pointer := 0
	right_pointer := len(height) - 1
	final_area := 0

	for left_pointer < right_pointer {
		current_area := smallest(height[left_pointer], height[right_pointer]) * (right_pointer - left_pointer)

		if current_area > final_area {
			final_area = current_area
		}

		if height[left_pointer] < height[right_pointer] {
			left_pointer++
		} else {
			right_pointer--
		}

	}
	return final_area
}

func smallest(left, right int) int {
	if left < right {
		return left
	}
	return right
}
