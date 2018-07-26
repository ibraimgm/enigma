package parts

// Plugboard represents a plugboard in an Enigma machine. The plugboard receives a signal
// and (maybe) change it to a different letter.
type Plugboard interface {
	Translate(input Signal) Signal
}
