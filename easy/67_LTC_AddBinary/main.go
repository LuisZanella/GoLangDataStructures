package main

/*
67. Add Binary
Given two binary strings a and b, return their sum as a binary string.

Example 1:

	Input: a = "11", b = "1"
	Output: "100"
	Example 2:

	Input: a = "1010", b = "1011"
	Output: "10101"

Constraints:

	1 <= a.length, b.length <= 104
	a and b consist only of '0' or '1' characters.
	Each string does not contain leading zeros except for the zero itself.

	func addBinary(a string, b string) string {
	    num1, _ := new(big.Int).SetString(a, 2)
		num2, _ := new(big.Int).SetString(b, 2)
		num1.Add(num1, num2)
		return num1.Text(2)
	}
*/
func addBinary(a string, b string) string {
	la := len(a) - 1
	lb := len(b) - 1
	if la < lb {
		return addBinary(b, a)
	}
	var carry int
	var sum int
	var result = make([]byte, la+1)
	for la >= 0 || lb >= 0 {
		if la >= 0 {
			sum = int(a[la]) + sum - '0'
		}
		if lb >= 0 {
			sum = int(b[lb]) + sum - '0'
		}

		result[la] = uint8(sum%2) + '0'
		carry = sum / 2
		la--
		lb--
		sum = carry
	}
	if sum == 1 {
		return string(append([]byte{'1'}, result...))
	}
	return string(result[:])
}
