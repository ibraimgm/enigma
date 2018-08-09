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
	Move(step int)
	Ring() int
	SetRing(value int)
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
		return CreateRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "R")
	case "II":
		return CreateRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "F")
	case "III":
		return CreateRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "W")
	case "IV":
		return CreateRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "K")
	case "V":
		return CreateRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "A")
	case "VI":
		return CreateRotor("JPGVOUMFYQBENHZRDKASXLICTW", "AN")
	case "VII":
		return CreateRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "AN")
	case "VIII":
		return CreateRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "AN")
	default:
		panic("Unrecognized rotor ID")
	}
}

type rotorImpl struct {
	position int
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

	return Rotor(&rotorImpl{position: 1, sequence: sequenceInt, notches: notchesInt})
}

func (r *rotorImpl) Window() rune {
	return intToChar(r.position)
}

func (r *rotorImpl) Move(step int) {
	newPos := r.position - 1 + step
	r.position = int(math.Mod(float64(newPos+26), 26)) + 1
}

func (r *rotorImpl) Ring() int {
	panic("Not implemented yet")
}

func (r *rotorImpl) SetRing(value int) {
	panic("Not implemented yet")
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
	from := r.position - 1 + int(input) - 1
	from = int(math.Mod(float64(from+26), 26))

	return Signal(r.sequence[from])
}

func (r *rotorImpl) Reverse(input Signal) Signal {
	in := int(input)

	for i, v := range r.sequence {
		if v == in {
			in = i + 1
			break
		}
	}

	return Signal(in)
}
