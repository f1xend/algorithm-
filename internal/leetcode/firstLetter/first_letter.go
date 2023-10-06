package firstletter

func repeatedCharacter(s string) byte {
	var l int
	if len(s) > 26 {
		l = 26
	} else {
		l = len(s) - 1
	}
	var m = make(map[rune]struct{}, l)

	for _, letter := range s {
		if _, ok := m[letter]; ok {
			return byte(letter)
		}
		m[letter] = struct{}{}
	}

	return 0
}
