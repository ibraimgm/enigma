package parts

// Keyboard defines a machine part that receives a character and returns a signal (number between 1 - 26)
// and a boolean (true if the conversion is ok).
type Keyboard interface {
	InputKey(key rune) (Signal, bool)
}

// DefaultKeyboard is the standard keyboard used in the machine.
// It accepts only the letters A-Z and a-z (converting to uppercase), and rejects any other character.
var DefaultKeyboard Keyboard = &keyboardImpl{}

type keyboardImpl struct{}

func (*keyboardImpl) InputKey(key rune) (Signal, bool) {
	chr := key

	if chr >= 'a' && chr <= 'z' {
		chr = chr - 32
	}

	s := charToInt(chr)

	return Signal(s), s != -1
}
