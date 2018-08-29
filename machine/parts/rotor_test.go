package parts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ibraimgm/enigma/machine/parts"
)

func TestRotorCreationID(t *testing.T) {
	rotors := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII"}

	for _, id := range rotors {
		r, _ := parts.GetRotor(id)
		assert.Equal(t, id, r.ID())
	}

	_, err := parts.GetRotor("XX")
	assert.EqualError(t, err, "unrecognized rotor ID: 'XX'")
}

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
	rotor, _ := parts.GetRotor("I")

	for _, test := range tests {
		rotor.Move(test.step)
		assert.Equal(t, test.window, rotor.Window())
	}
}

func TestRotorIVWindowAndMove(t *testing.T) {
	rotor, _ := parts.GetRotor("IV")
	assert.Equal(t, 'A', rotor.Window())

	rotor.SetWindow('C')
	assert.Equal(t, 'C', rotor.Window())

	rotor.Move(3)
	assert.Equal(t, 'F', rotor.Window())

	rotor.SetWindow('Y')
	rotor.Move(3)
	assert.Equal(t, 'B', rotor.Window())
}

func TestRotorISimpleStep(t *testing.T) {
	tests := []rotorStepTable{
		{1, 'A', 'B', false, false, 1, 10}, // A -> J
		{1, 'B', 'C', false, false, 2, 4},  // B -> D
		{1, 'C', 'D', false, false, 2, 9},  // B -> I
		{1, 'D', 'E', false, false, 1, 8},  // A -> H
	}

	r, _ := parts.GetRotor("I")
	rotorStepTableRunner(t, r, tests)
}

func TestRotorIWeirdStep(t *testing.T) {
	tests := []rotorStepTable{
		{1, 'A', 'B', false, false, 1, 10},  // A -> J
		{2, 'B', 'D', false, false, 2, 9},   // B -> I
		{1, 'D', 'E', false, false, 2, 3},   // B -> C
		{-5, 'E', 'Z', false, false, 1, 11}, // A -> K
	}

	r, _ := parts.GetRotor("I")
	rotorStepTableRunner(t, r, tests)
}

func TestRotorIINotch(t *testing.T) {
	tests := []rotorStepTable{
		{2, 'A', 'C', false, false, 17, 24}, // Q -> X
		{1, 'C', 'D', false, false, 21, 19}, // U -> S
		{1, 'D', 'E', false, true, 5, 20},   // E -> T
		{1, 'E', 'F', true, false, 5, 23},   // E -> W
		{1, 'F', 'G', false, false, 14, 8},  // N -> H
	}

	r, _ := parts.GetRotor("II")
	rotorStepTableRunner(t, r, tests)
}

func TestRotorVINotch(t *testing.T) {
	tests := []rotorStepTable{
		{1, 'A', 'B', false, false, 5, 20},  // E -> T
		{1, 'B', 'C', false, false, 14, 16}, // N -> P
		{1, 'C', 'D', false, false, 9, 2},   // I -> B
		{9, 'D', 'M', false, true, 7, 15},   // G -> O
		{1, 'M', 'N', true, false, 13, 10},  // M -> J
		{12, 'N', 'Z', false, true, 1, 24},  // A -> X
		{1, 'Z', 'A', true, false, 19, 1},   // S -> A
	}

	r, _ := parts.GetRotor("VI")
	rotorStepTableRunner(t, r, tests)
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
	rotor3, _ := parts.GetRotor("III")
	rotor2, _ := parts.GetRotor("II")
	rotor1, _ := parts.GetRotor("I")

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

func TestRotorScrambleWithRing(t *testing.T) {
	rotor3, _ := parts.GetRotor("III")
	rotor2, _ := parts.GetRotor("II")
	rotor1, _ := parts.GetRotor("I")

	assert.Equal(t, 'A', rotor3.Ring())
	rotor3.SetRing('B')
	assert.Equal(t, 'B', rotor3.Ring())

	tests := []rotorScrambleTable{
		{rotor3, true, 7, 13},  // G -> M
		{rotor2, true, 13, 23}, // M -> W
		{rotor1, true, 23, 2},  // W -> B
		//reflector: B -> R
		{rotor1, false, 18, 24}, // R -> X
		{rotor2, false, 24, 9},  // X -> I
		{rotor3, false, 9, 5},   // I -> E
	}

	rotorScrambleTableRunner(t, tests)
}

func TestRotorScrambleMultipleRing(t *testing.T) {
	rotor3, _ := parts.GetRotor("V")
	rotor2, _ := parts.GetRotor("VII")
	rotor1, _ := parts.GetRotor("VIII")

	rotor1.SetWindow('I')
	rotor2.SetWindow('R')
	rotor3.SetWindow('M')

	rotor1.SetRing('C')
	rotor2.SetRing('D')
	rotor3.SetRing('R')

	tests := []rotorScrambleTable{
		{rotor3, true, 2, 11},  // B -> K
		{rotor2, true, 11, 16}, // K -> P
		{rotor1, true, 16, 3},  // P -> C
		//reflector: C -> U
		{rotor1, false, 21, 11}, // U -> K
		{rotor2, false, 11, 22}, // K -> V
		{rotor3, false, 22, 26}, // V -> Z
	}

	rotorScrambleTableRunner(t, tests)
}
