package main

import "fmt"

func main() {
	s := "sas"
	t := "assa"
	isAnagram(s, t)

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := [26]int{} // 26 for the alphabet chars
	mapT := [26]int{}

	for i := 0; i < len(s); i++ {
		fmt.Println(mapS)
		fmt.Println(mapT)
		mapS[s[i]-'a']++
		mapT[t[i]-'a']++
	}
	fmt.Println(mapS == mapT)
	return mapS == mapT
}

// this solution is a binary way to figure it out
//we place a 1 if the char value is seen, and stays 0 if not see
//at the end both arrays if similar should look the same
//if they dont then it is not an anagram
