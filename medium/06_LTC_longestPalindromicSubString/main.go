package main

func longestPalindrome(s string) string {
	var low int
	var end int
	n := len(s)
	for i := 0; i < n; i++ {
		r := i
		low, end = findPalindromeMinAndMaxIndexes(low, end, i, r, s)
		r = i + 1
		low, end = findPalindromeMinAndMaxIndexes(low, end, i, r, s)
	}
	if low == 0 && end == 0 && n > 0 {
		return s[low : end+1]
	}
	return s[low : low+end]
}

func findPalindromeMinAndMaxIndexes(init, end, i, r int, w string) (int, int) {
	l := i - 1
	for l >= 0 && r < len(w) && w[l] == w[r] {
		slice := r - l + 1
		if slice > end {
			init = l
			end = slice
		}
		l--
		r++
	}
	return init, end
}
