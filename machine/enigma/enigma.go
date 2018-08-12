package enigma

import (
	"github.com/ibraimgm/enigma/machine/parts"
)

// Enigma is a interface describing a generic 3-rotor enigma machine.
// Once built, you cannot change the machine parts, but can configure the window settings and the ring settings
type Enigma interface {
	Window() string
	SetWindow(settings string)
	Ring() string
	SetRing(settings string)
	Configure(ringSetting, windowSetting string)
	Encode(input rune) (rune, bool)
	EncodeMessage(message string, blockSize uint) string
}

type enigmaImpl struct {
	keyboard   parts.Keyboard
	plugboard  parts.Plugboard
	slow       parts.Rotor
	middle     parts.Rotor
	fast       parts.Rotor
	reflector  parts.Reflector
	lightboard parts.Lightboard
}

// WithDefaults builds a new enigma machine, with the rotors III, II and I (from slow to fast), using the "B" reflector
// with both window settings and ring settings to 'AAA'.
func WithDefaults() Enigma {
	return Assemble(
		parts.DefaultKeyboard,
		parts.NoPlugboard,
		parts.GetRotor("III"),
		parts.GetRotor("II"),
		parts.GetRotor("I"),
		parts.Reflectors["B"],
		parts.DefaultLightboard)
}

// WithRotors build a new enigma machine, with the specified rotors and reflector.
func WithRotors(slow, middle, fast, reflector string) Enigma {
	return Assemble(
		parts.DefaultKeyboard,
		parts.NoPlugboard,
		parts.GetRotor(slow),
		parts.GetRotor(middle),
		parts.GetRotor(fast),
		parts.Reflectors[reflector],
		parts.DefaultLightboard,
	)
}

// WithConfig builds a new enigma machine, with default rotors and reflector, using the specified settings.
func WithConfig(ringSetting, windowSetting string) Enigma {
	e := WithDefaults()
	e.Configure(ringSetting, windowSetting)

	return e
}

// Assemble builds a new enigma machine, with default config and all parts specified.
// This is the only way to create a machine with a different keyboard, lightboard or plugboard
func Assemble(keyboard parts.Keyboard, plugboard parts.Plugboard, slow, middle, fast parts.Rotor, reflector parts.Reflector, lightboard parts.Lightboard) Enigma {
	enigma := &enigmaImpl{keyboard, plugboard, slow, middle, fast, reflector, lightboard}
	enigma.SetWindow("AAA")
	enigma.SetRing("AAA")

	return Enigma(enigma)
}

func (e *enigmaImpl) Window() string {
	return string([]rune{
		e.slow.Window(),
		e.middle.Window(),
		e.fast.Window(),
	})
}

func (e *enigmaImpl) SetWindow(settings string) {
	var runes []rune

	if settings == "" {
		runes = []rune{'A', 'A', 'A'}
	} else {
		runes = []rune(settings)
	}

	e.slow.SetWindow(runes[0])
	e.middle.SetWindow(runes[1])
	e.fast.SetWindow(runes[2])
}

func (e *enigmaImpl) Ring() string {
	return string([]rune{
		e.slow.Ring(),
		e.middle.Ring(),
		e.fast.Ring(),
	})
}

func (e *enigmaImpl) SetRing(settings string) {
	var runes []rune

	if settings == "" {
		runes = []rune{'A', 'A', 'A'}
	} else {
		runes = []rune(settings)
	}

	e.slow.SetRing(runes[0])
	e.middle.SetRing(runes[1])
	e.fast.SetRing(runes[2])
}

func (e *enigmaImpl) Configure(ringSetting, windowSetting string) {
	e.SetRing(ringSetting)
	e.SetWindow(windowSetting)
}

func (e *enigmaImpl) Encode(input rune) (rune, bool) {
	// only run on valid signals
	signal, ok := e.keyboard.InputKey(input)
	if !ok {
		return input, false
	}

	// stepping
	if e.fast.IsNotched() {
		if e.middle.IsNotched() {
			e.slow.Move(1)
		}

		e.middle.Move(1)
	} else if e.middle.IsNotched() {
		e.middle.Move(1)
		e.slow.Move(1)
	}

	e.fast.Move(1)

	// signal flow
	signal = e.plugboard.Translate(signal)
	signal = e.fast.Scramble(signal)
	signal = e.middle.Scramble(signal)
	signal = e.slow.Scramble(signal)
	signal = e.reflector.Reflect(signal)
	signal = e.slow.Reverse(signal)
	signal = e.middle.Reverse(signal)
	signal = e.fast.Reverse(signal)
	signal = e.plugboard.Translate(signal)
	return e.lightboard.Light(signal), true
}

func (e *enigmaImpl) EncodeMessage(message string, blockSize uint) string {
	var currSize uint
	var currMsg string

	for _, c := range message {
		encoded, ok := e.Encode(c)
		if !ok {
			continue
		}

		if blockSize > 0 && currSize == blockSize {
			currMsg = currMsg + " "
			currSize = 0
		}

		currMsg = currMsg + string(encoded)
		currSize++
	}

	return currMsg
}
