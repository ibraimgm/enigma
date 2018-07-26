package parts

// Lightboard receives a signal and converts it to a valid character.
// It is the opposite of the Keyboard, and a custom Keyboard implementatio will most likely also provide a custom
// Lightboard implementation.
type Lightboard interface {
	Light(input Signal) rune
}
