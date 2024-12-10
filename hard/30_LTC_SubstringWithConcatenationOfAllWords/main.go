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

func createHeatMap(words []string) map[string]int {
	heatMap := make(map[string]int)
	for _, word := range words {
		heatMap[word]++
	}
	return heatMap
}

func FindSubstring(s string, words []string) (indexSubstrings []int) {
	if s == "" || len(words) == 0 || len(words[0]) == 0 {
		return indexSubstrings
	}
	subWordLen := len(words[0])
	subStringMinSize := subWordLen * len(words)

	if len(s) < subStringMinSize {
		return indexSubstrings
	}

	wordFrequency := createHeatMap(words)
	seenWords := make(map[string]int)
	windowSize := len(s) - subStringMinSize
	for i := 0; i <= windowSize; i++ {
		j := 0
		for seenWord := range seenWords {
			delete(seenWords, seenWord)
		}
		for j < len(words) {
			start := i + (j * subWordLen)
			word := s[start : start+subWordLen]
			if freq, exists := wordFrequency[word]; exists {
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
