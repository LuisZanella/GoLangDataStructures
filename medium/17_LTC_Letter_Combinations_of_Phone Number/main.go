/**

1(phone)|2(abc)		|3(def)
4(ghi)	|5(jkl)		|6(mno)
7(pqrs)	|8(tuv)		|9(wxyz)
*(+)	|0(space)	|#(caps)

*/
//Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could
//represent. Return the answer in any order.
//
//A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1
//does not map to any letters.
//

package main

import (
	"fmt"
)

var phoneMap = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var result []string
	backTracking(0, "", &digits, &result)
	return result
}

func backTracking(index int, combination string, digits *string, result *[]string) {
	if index == len(*digits) {
		*result = append(*result, combination)
		return
	}
	letters := phoneMap[rune((*digits)[index])]
	for _, letter := range letters {
		backTracking(index+1, combination+letter, digits, result)
	}
}

func main() {
	digits := "23"
	combinations := letterCombinations(digits)
	fmt.Println(combinations)
}
