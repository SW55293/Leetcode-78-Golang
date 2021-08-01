package main

func main() {
	arr := []int{2, 7, 89, 1, 6}
	target := 8
	twoSum(arr, target)
}

func twoSum(arr []int, target int) []int {

	copy := make(map[int]int)

	for key, value := range arr {

		if index, ok := copy[target-value]; ok {
			return []int{index, key}
		}
		copy[value] = key
	}

	return []int{0, 0}
}
