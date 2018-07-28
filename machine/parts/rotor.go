package parts

// Rotor also known as 'scrambler' is the main piece that controls how the text
// is transformed in an Enigma machine. It has a 'window' that shows it's current position,
// and is able to convert the signal both from the 'left' (Scramble method) as well as from
// the 'right' (reverse method), after the signal is reflected by the reflector.
type Rotor interface {
	Window() Signal
	Move() bool
	Scramble(input Signal) Signal
	Reverse(input Signal) Signal
}
