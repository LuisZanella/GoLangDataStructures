package main

func isUnique(s string) bool {
	var asciiArray = [128]bool{}
	for _, i := range s {
		if i > 128 {
			return false
		}
		if asciiArray[i] == true {
			return false
		}
		asciiArray[i] = true
	}
	return true
}
