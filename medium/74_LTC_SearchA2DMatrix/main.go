package main

/*
You are given an m x n integer matrix matrix with the following two properties:

 - Each row is sorted in non-decreasing order.
 - The first integer of each row is greater than the last integer of the previous row.

Given an integer target, return true if target is in matrix or false otherwise.
You must write a solution in O(log(m * n)) time complexity.
*/

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, (m*n)-1
	for l <= r {
		m := l + (r-l)/2
		row := m / 2
		col := m % 2

		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			r = m - 1
		} else {
			l = m + 1
		}

	}
	return false
}

func main() {

}
