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
*/

import (
	"bytes"
	"fmt"
)

func findCommonPrefix(word string, defaultPrefix string) string {
	var prefix bytes.Buffer
	var newLength = len(word)
	var prefixLength = len(defaultPrefix)
	if newLength > prefixLength {
		newLength = prefixLength
	}
	for i := 0; i < newLength; i++ {
		if word[i] != defaultPrefix[i] {
			break
		}
		prefix.WriteByte(defaultPrefix[i])
	}
	return prefix.String()
}
func mapCommonPrefix(words []string) string {
	var arraySize = len(words)
	var prefix = words[0]
	for i := 1; i < arraySize; i++ {
		prefix = findCommonPrefix(words[i], prefix)
	}
	return prefix
}

func main() {
	var e = []string{"apple", "app", "appleWatch"}
	fmt.Println(mapCommonPrefix(e))
	e = []string{"dog", "racecar", "car"}
	fmt.Println(mapCommonPrefix(e))
	e = []string{"flower", "flow", "flight"}
	fmt.Println(mapCommonPrefix(e))
}

/**
	func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    slices.Sort(strs)
    last := strs[len(strs) - 1]
    first := strs[0]

    prefixLen := 0
    shortestStrLen := len(last)
    if len(first) < shortestStrLen {
        shortestStrLen = len(first)
    }

    for i := 0; i < shortestStrLen; i++ {
        if last[i] != first[i] {
            break
        }

        prefixLen++
    }

    return first[:prefixLen]
}


VS

func longestCommonPrefix(strs []string) string {

    ret := ""
    limit := 0
    if len(strs) == 1 {
        return strs[0]
    }
    sort.Strings(strs)
    start := strs[0]
    end := strs[len(strs) -1]
    if len(start) < len(end) {
        limit = len(start)
    } else {
        limit = len(end)
    }
    for i := 0; i < limit && start[i] == end[i]; i++ {
        ret += string(start[i])
        // fmt.Printf("short = %c\n", start[i])
    }
    return ret
}
*/
