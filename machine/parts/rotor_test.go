package parts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ibraimgm/enigma/machine/parts"
)

type rotorStepTable struct {
	step          int
	windowBefore  rune
	windowAfter   rune
	notchedBefore bool
	notchedAfter  bool
	toScramble    parts.Signal
	scrambled     parts.Signal
}

func rotorStepTableRunner(t *testing.T, rotor parts.Rotor, tests []rotorStepTable) {
	for _, test := range tests {
		assert.Equal(t, test.windowBefore, rotor.Window())
		assert.Equal(t, test.notchedBefore, rotor.IsNotched())
		rotor.Move(test.step)
		assert.Equal(t, test.windowAfter, rotor.Window())
		assert.Equal(t, test.notchedAfter, rotor.IsNotched())

		scrambled := rotor.Scramble(test.toScramble)
		assert.Equal(t, test.scrambled, scrambled)
	}
}

func TestRotorIMove(t *testing.T) {
	var tests = []struct {
		step   int
		window rune
	}{
		{1, 'B'},
		{2, 'D'},
		{-1, 'C'},
		{-3, 'Z'},
		{1, 'A'},
		{15, 'P'},
		{9, 'Y'},
		{2, 'A'},
		{-3, 'X'},
		{4, 'B'},
	}
	rotor := parts.GetRotor("I")

	for _, test := range tests {
		rotor.Move(test.step)
		assert.Equal(t, test.window, rotor.Window())
	}
}

func TestRotorISimpleStep(t *testing.T) {
	tests := []rotorStepTable{
		{1, 'A', 'B', false, false, 1, 11}, // A -> K
		{1, 'B', 'C', false, false, 2, 6},  // B -> F
		{1, 'C', 'D', false, false, 2, 12}, // B -> L
		{1, 'D', 'E', false, false, 1, 12}, // A -> L
	}

	rotorStepTableRunner(t, parts.GetRotor("I"), tests)
}

func TestRotorIWeirdStep(t *testing.T) {
	tests := []rotorStepTable{
		{1, 'A', 'B', false, false, 1, 11},  // A -> K
		{2, 'B', 'D', false, false, 2, 12},  // B -> L
		{1, 'D', 'E', false, false, 2, 7},   // B -> G
		{-5, 'E', 'Z', false, false, 1, 10}, // A -> J
	}

	rotorStepTableRunner(t, parts.GetRotor("I"), tests)
}

func TestRotorIINotch(t *testing.T) {
	tests := []rotorStepTable{
		{2, 'A', 'C', false, false, 17, 26}, // Q -> Z
		{1, 'C', 'D', false, false, 21, 22}, // U -> V
		{1, 'D', 'E', false, false, 5, 24},  // E -> X
		{1, 'E', 'F', false, true, 5, 2},    // E -> B
		{1, 'F', 'G', true, false, 14, 14},  // N -> N
	}

	rotorStepTableRunner(t, parts.GetRotor("II"), tests)
}

type rotorScrambleTable struct {
	rotor       parts.Rotor
	rightToLeft bool
	in          int
	out         int
}

func rotorScrambleTableRunner(t *testing.T, tests []rotorScrambleTable) {
	for _, test := range tests {
		rotor := test.rotor

		if test.rightToLeft {
			assert.Equal(t, parts.Signal(test.out), rotor.Scramble(parts.Signal(test.in)))
		} else {
			assert.Equal(t, parts.Signal(test.out), rotor.Reverse(parts.Signal(test.in)))
		}
	}
}

func TestRotorScramble(t *testing.T) {
	rotor3 := parts.GetRotor("III")
	rotor2 := parts.GetRotor("II")
	rotor1 := parts.GetRotor("I")

	tests := []rotorScrambleTable{
		{rotor3, true, 7, 3}, // G -> C
		{rotor2, true, 3, 4}, // C -> D
		{rotor1, true, 4, 6}, // D -> F
		//reflector: F -> S
		{rotor1, false, 19, 19}, // S -> S
		{rotor2, false, 19, 5},  // S -> E
		{rotor3, false, 5, 16},  // E -> P
	}

	rotorScrambleTableRunner(t, tests)
}
