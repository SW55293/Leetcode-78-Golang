package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 8}
	k := 6
	subarraySum(s, k)
}

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	mapp := make(map[int]int)
	mapp[0] = 1
	fmt.Println(mapp)
	longestSub := 0
	sum := 0

	for _, value := range nums {
		sum += value
		// fmt.Println("Sum = ", sum) //sum of values up to the index we are at
		// fmt.Println("value = ", value) // values in the int array

		longestSub += mapp[sum-k]
		// fmt.Println("mapp[sum-k] = ", mapp[sum-k])

		mapp[sum]++
		fmt.Println("mapp[sum]++ = ", mapp) //just adds the sum to the map and increments counter for the occurrence
	}
	return longestSub
}

/*
func subarraySum(nums []int, k int) int {

	counter := 0
	sum := 0 //keeps track of the largest subarray sum total so far. That is <= the value of k

	newMap := make(map[int]int) //1st int =sum of current and all that came before it, 2nd int =
	newMap[0] = 1

	for x := 0; x < len(nums); x++ {

		sum = sum + nums[x] //sum = sum + nums[x=0,1,2,3,4..]
		fmt.Println(sum - k)
		if val, ok := newMap[sum-k]; ok {

			//fmt.Println("val = ", val)

			counter = counter + val //counter = counter + val

		}

		newMap[sum] += 1
		//fmt.Println("map = ", newMap)

	}
	//fmt.Println("counter = ", counter)
	return counter

}
*/
