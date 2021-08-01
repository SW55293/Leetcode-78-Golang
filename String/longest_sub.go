package main

/*
a = 97
b = 98
c = 99
r = 114

s = 0    1   2   3    4   5
	97  98  99  114  98  99
*/
func main() {
	s := "abcrbc"
	lengthOfLongestSubstring(s)
}

//two pointer technique
func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	longest_sub := 0

	visited := make(map[rune]int)

	for index, value := range s {

		end = index

		if x, found := visited[value]; found {

			if x+1 > start {

				start = x + 1

			}
		}
		if longest_sub < end-start+1 {
			longest_sub++
		}
		visited[value] = end
	}

	return longest_sub
}

/*
	fmt.Println("current max = ", current_max)
		fmt.Println("start = ", start)
		fmt.Println("key = ", key)
		fmt.Println(longest_sub)
*/
