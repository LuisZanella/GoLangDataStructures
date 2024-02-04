package main

func RotateMatrixNxN(m [][]int) [][]int {
	ms := len(m)
	for layer := 0; layer < ms/2; layer++ {
		f := layer
		l := ms - 1 - layer
		for i := f; i < l; i++ {
			of := i - f
			top := m[f][i]
			// left -> top
			m[f][i] = m[l-of][f]
			// bottom -> left
			m[l-of][f] = m[l][l-of]
			// right -> bottom
			m[l][l-of] = m[i][l]
			// top -> right
			m[i][l] = top
		}
	}
	return m
}
