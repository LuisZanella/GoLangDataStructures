package main

func isPalindromePermutation(s string) bool {
	var odd int
	mapper := map[int32]int{}
	for _, val := range s {
		if val >= 'A' && val <= 'Z' {
			val = val + ('a' - 'A')
		}
		if val > 'z' || val < 'a' {
			continue
		}
		mapper[val]++
		c := mapper[val]
		if c%2 == 1 {
			odd++
		} else {
			odd--
		}

	}
	return odd <= 1
}
