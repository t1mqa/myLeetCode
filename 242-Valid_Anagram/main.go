package main

import "fmt"

func isAnagram(s string, t string) bool {
	// Check if strings have same size
	if len(s) != len(t) {
		return false
	}

	// mega effective RAM solution
	counter := [26]int{}

	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("rat", "car"))
}
