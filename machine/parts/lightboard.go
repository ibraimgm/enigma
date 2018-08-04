package parts

// Lightboard receives a signal and converts it to a valid character.
// It is the opposite of the Keyboard, and a custom Keyboard implementatio will most likely also provide a custom
// Lightboard implementation.
type Lightboard interface {
	Light(input Signal) rune
}

// DefaultLightboard converts a signal into an uppercase letter. A = 1, Z = 26
var DefaultLightboard Lightboard = &lightboardImpl{}

type lightboardImpl struct{}

func (*lightboardImpl) Light(input Signal) rune {
	return intToChar(int(input))
}
