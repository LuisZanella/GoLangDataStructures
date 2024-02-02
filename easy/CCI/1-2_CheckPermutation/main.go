package main

func arePermutation(v []string) bool {
	w1 := v[0]
	w2 := v[1]
	mapper := map[int32]int{}
	// We can use a number array instead if we want
	for _, c := range w1 {
		mapper[c]++
	}
	for _, c := range w2 {
		val, ok := mapper[c]
		if !ok || val <= 0 {
			return false
		}
		mapper[c]--
	}
	return true
}
