package parts

// Signal is the value that actually traverses the enigma machine.
// It is generated by a Keyboard implementation, and generally is a number between 1 and 26, representing each letter in
// the alphabet.
type Signal int

func charToInt(c rune) int {
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 1
	}

	return -1
}

func intToChar(i int) rune {
	if i >= 1 && i <= 26 {
		return rune(i - 1 + 'A')
	}

	return 0
}

func fixAlpha(i int) int {
	if i <= 0 {
		return i + 26
	} else if i > 26 {
		return i - 26
	}

	return i
}
