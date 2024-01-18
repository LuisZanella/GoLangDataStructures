package main
/**
	Write a function to find the longest common prefix string amongst an array of strings.

	If there is no common prefix, return an empty string "".

	

	Example 1:

	Input: strs = ["flower","flow","flight"]
	Output: "fl"
	Example 2:

	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.
	

	Constraints:

	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lowercase English letters.
**/
import ( "fmt")

func commonPrefix(words []string) string {
	if len(words) == 0 {
		return ""
	}
	var firstWordLength int = len(words[0]) - 1
	var commonPrefix = make(string, firstWordLength)
	var counter = 0
	for i,value range words {
		if words[0][i] == value {
			commonPrefix[i] = value
		}
		
	} 
	for counter < firstWordLength && firstWordLength <  {
		counter++
	}
	return commonPrefix
}

func main () {
	fmt.Println("Testing")
}

