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

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	// Mapping of digits to corresponding letters
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			result = append(result, path)
			return
		}

		letters := digitToLetters[digits[index]]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, path+string(letters[i]))
		}
	}

	backtrack(0, "")
	return result
}

func main() {
	digits := "23"
	combinations := letterCombinations(digits)
	fmt.Println(combinations)
}
