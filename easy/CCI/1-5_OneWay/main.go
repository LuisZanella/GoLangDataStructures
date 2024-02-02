package main

func Oneway(w []string) bool {
	var diffs int
	w1 := len(w[0])
	w2 := len(w[1])

	if w1 > w2 {
		return Oneway([]string{w[1], w[0]})
	}
	if w2-w1 > 1 {
		return false
	}
	if w1 != w2 {
		diffs++
	}

	for i, _ := range w[0] {
		if w[0][i] != w[1][i] {
			diffs++
		}
		if diffs > 1 {
			return false
		}
	}
	return true
}
