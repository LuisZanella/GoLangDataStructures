package main

func urlLify(s string, l int) string {
	new_str := []byte(s)
	sSize := len(s) - 1
	l = l - 1
	for i := sSize; i >= 0; i-- {
		if s[l] != 32 {
			new_str[i] = s[l]
			l--
			continue
		}
		new_str[i] = 48
		i--
		new_str[i] = 50
		i--
		new_str[i] = 37
		l--
	}
	return string(new_str)
}
