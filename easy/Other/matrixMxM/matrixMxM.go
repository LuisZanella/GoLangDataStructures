package main

//TODO: FIX This is rotating wrong -90 degrees
import "fmt"

func rotateMatrix90(matrix [][]int) [][]int {
	var matrixSize = len(matrix)
	var newX int = 0
	var newY int = 0
	newMatrix := make([][]int, matrixSize)
	for y := matrixSize - 1; y >= 0; y-- {
		newMatrix[newY] = make([]int, matrixSize)
		newX = 0
		for x := 0; x < matrixSize; x++ {
			newMatrix[newY][newX] = matrix[x][y]
			newX++
		}
		newY++
	}
	return newMatrix
}

func main() {
	var matrix [][]int
	matrix = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	result := rotateMatrix90(matrix)
	fmt.Println(result)

	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	result = rotateMatrix90(matrix)
	fmt.Println(result)
}
