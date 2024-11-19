package main

import (
	"fmt"
)

func createHash(rawStr string) [26]byte {
	// hash - array of length 26 (alphabet)
	// Each index is letter - a=0, z=26
	hash := [26]byte{}
	for _, symb := range rawStr {
		hash[symb-'a']++
	}
	return hash
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	var groups [][]string
	strsMap := make(map[[26]byte][]int)
	// strsMap - map, where keys = hashes, and values = indexes

	// strs = ["abc", "def", "bca"]
	for ind, str := range strs {
		hash := createHash(str)
		strsMap[hash] = append(strsMap[hash], ind)
	}

	// now lets built groups, using strsMap
	for _, value := range strsMap {
		// values - list of indexes
		// convert indexes to group of words
		var words []string

		for _, ind := range value {
			words = append(words, strs[ind])
		}

		// insert group of words inside groups
		groups = append(groups, words)
	}

	return groups
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
