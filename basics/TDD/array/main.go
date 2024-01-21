package main

import "errors"

func SumArray(numbers []int) int {
	var total int
	for _, value := range numbers {
		total += value
	}
	return total
}

func SumTwoArraysPerPosition(n1 []int, n2 []int) ([]int, error) {
	if len(n1) != len(n2) {
		return []int{}, errors.New("Elements doesn't have same length")
	}
	for i, value := range n1 {
		n2[i] += value
	}
	return n2, nil
}
