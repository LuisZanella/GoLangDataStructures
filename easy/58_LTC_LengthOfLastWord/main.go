package main

/*
58. Length of Last Word
Easy
Topics
Companies
Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal
substring
 consisting of non-space characters only.

Example 1:

	Input: s = "Hello World"
	Output: 5
	Explanation: The last word is "World" with length 5.
	Example 2:

	Input: s = "   fly me   to   the moon  "
	Output: 4
	Explanation: The last word is "moon" with length 4.
	Example 3:

	Input: s = "luffy is still joyboy"
	Output: 6
	Explanation: The last word is "joyboy" with length 6.


Constraints:

	1 <= s.length <= 104
	s consists of only English letters and spaces ' '.
	There will be at least one word in s.

*/
//YOOOOO 10/10
func lengthOfLastWord(s string) int {
	size := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			size++
		} else {
			if size == -1 {
				continue
			} else {
				return size + 1
			}
		}
	}
	return size + 1
}
