package parts

import (
	"math"
)

// Rotor also known as 'scrambler' is the main piece that controls how the text
// is transformed in an Enigma machine. It has a 'window' that shows it's current position,
// and is able to convert the signal both from the 'left' (Scramble method) as well as from
// the 'right' (reverse method), after the signal is reflected by the reflector.
type Rotor interface {
	Window() rune
	SetWindow(value rune)
	Move(step int)
	Ring() rune
	SetRing(value rune)
	IsNotched() bool
	Scramble(input Signal) Signal
	Reverse(input Signal) Signal
}

// GetRotor returns a default implementation of one of the historical rotors.
// The id must be one of the roman numerals, from I to VIII.
// Each call to GetRotor returns a new instance.
func GetRotor(id string) Rotor {

	switch id {
	case "I":
		return CreateRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "Q")
	case "II":
		return CreateRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "E")
	case "III":
		return CreateRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "V")
	case "IV":
		return CreateRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "J")
	case "V":
		return CreateRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "Z")
	case "VI":
		return CreateRotor("JPGVOUMFYQBENHZRDKASXLICTW", "ZM")
	case "VII":
		return CreateRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "ZM")
	case "VIII":
		return CreateRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "ZM")
	default:
		panic("Unrecognized rotor ID")
	}
}

type rotorImpl struct {
	position int
	ring     int
	sequence []int
	notches  []int
}

// CreateRotor creates a new rotor, with the specified sequence (a 26 letter string)
// and the specified notches (a string with one or more chracters)
func CreateRotor(sequence string, notches string) Rotor {
	notchesRunes := []rune(notches)
	notchesInt := make([]int, len(notchesRunes))

	for i := range notchesRunes {
		notchesInt[i] = charToInt(notchesRunes[i])
	}

	sequenceRunes := []rune(sequence)
	sequenceInt := make([]int, len(sequenceRunes))

	for i := range sequenceRunes {
		sequenceInt[i] = charToInt(sequenceRunes[i])
	}

	return Rotor(&rotorImpl{position: 1, ring: 1, sequence: sequenceInt, notches: notchesInt})
}

func (r *rotorImpl) Window() rune {
	return intToChar(r.position)
}

func (r *rotorImpl) SetWindow(value rune) {
	r.position = fixAlpha(charToInt(value))
}

func (r *rotorImpl) Move(step int) {
	newPos := r.position - 1 + step
	r.position = int(math.Mod(float64(newPos+26), 26)) + 1
}

func (r *rotorImpl) Ring() rune {
	return intToChar(r.ring)
}

func (r *rotorImpl) SetRing(value rune) {
	r.ring = fixAlpha(charToInt(value))
}

func (r *rotorImpl) IsNotched() bool {

	for _, v := range r.notches {
		if r.position == v {
			return true
		}
	}

	return false
}

func (r *rotorImpl) Scramble(input Signal) Signal {
	from := fixAlpha(int(input) - r.ring + r.position)
	from = r.sequence[from-1]
	from = fixAlpha(from + r.ring - r.position)

	return Signal(from)
}

func (r *rotorImpl) Reverse(input Signal) Signal {
	from := fixAlpha(int(input) - r.ring + r.position)

	for i, v := range r.sequence {
		if v == from {
			from = i + 1
			break
		}
	}

	from = fixAlpha(from + r.ring - r.position)
	return Signal(from)
}
