package main

import "fmt"

//You are given a string s and an array of strings words.
//All the strings of words are of the same length.
//
//A concatenated string is a string that exactly
//contains all the strings of any permutation of
//words concatenated.
//
//For example, if words = ["ab","cd","ef"], then
//"abcdef", "abefcd", "cdabef", "cdefab", "efabcd",
//and "efcdab" are all concatenated strings. "acdbef"
//is not a concatenated string because it is not the
//concatenation of any permutation of words.
//Return an array of the starting indices of all the
//concatenated substrings in s. You can return the
//answer in any order.

var permutationsPosition map[string]int

func generatePermutationsRecursive(str string, strIndex int) {
	var permutations []string
	if len(str) == 1 {
		permutationsPosition[str] = strIndex
		return append(permutations, str)
	}

	for i, c := range str {
		remaining := str[:i] + str[i+1:]
		subPermutations := generatePermutationsRecursive(remaining, strIndex)

		for _, p := range subPermutations {
			permutationsPosition[string(c)+p] = strIndex
			permutations = append(permutations, string(c)+p)
		}
	}

	return permutations
}

func findSubstring(s string, words []string) (result []string) {
	var startIndexes []int
	for i, word := range words {
		generatePermutationsRecursive(word, i)
	}
	substringSize := len(words[0]) - 1
	var left int
	right := len(words[0]) - 1
	for i := 0; i < substringSize; i++ {
		subString := s[left:right]
		if permutationsPosition[subString] != 0 {
			startIndexes = append(startIndexes, permutationsPosition[subString]-1)
		}
		left += substringSize
		right += substringSize
	}
	return startIndexes
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
