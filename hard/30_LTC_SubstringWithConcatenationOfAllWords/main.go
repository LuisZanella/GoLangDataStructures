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

func FindSubstring(s string, words []string) (indexSubstrings []int) {
	/*
		validate input empty cases
		generate HeatMap
		createWindowSlicing
			verify if window has the same as Heatmap
			if yes, add to indexSubstrings
		if no, continue next window
	*/
	if s == "" || len(words) == 0 || len(words[0]) == 0 {
		return indexSubstrings
	}
	subWordLen := len(words[0])
	subStringMinSize := subWordLen * len(words)

	if len(s) < subStringMinSize {
		return indexSubstrings
	}

	heatMap := make(map[string]int)
	for _, word := range words {
		heatMap[word]++
	}
	seenWords := make(map[string]int)
	windowSize := len(s) - subStringMinSize
	j := 0
	for i := 0; i <= windowSize; i++ {
		for seenWord := range seenWords {
			delete(seenWords, seenWord)
		}
		for j < len(words) {
			start := i + (j * subWordLen)
			word := s[start : start+subWordLen]
			if freq, exists := heatMap[word]; exists {
				seenWords[word]++
				if seenWords[word] > freq {
					break
				}
			} else {
				break
			}
			j++
		}
		if j == len(words) {
			indexSubstrings = append(indexSubstrings, i)
		}
	}
	return indexSubstrings
}

func main() {
	fmt.Println(FindSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
