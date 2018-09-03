//Package enigma provides a quick and easy way to create and assemble a new enigma machine instance.
package enigma

import (
	"errors"

	"github.com/ibraimgm/enigma/machine/parts"
)

// Enigma is a interface describing a generic 3-rotor enigma machine.
// Once built, you cannot change the machine parts, but can configure the window settings and the ring settings
type Enigma interface {
	Reflector() string
	Slow() string
	Middle() string
	Fast() string
	Window() string
	SetWindow(settings string) error
	Ring() string
	SetRing(settings string) error
	Configure(ringSetting, windowSetting string) error
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
	slow, _ := parts.GetRotor("III")
	middle, _ := parts.GetRotor("II")
	fast, _ := parts.GetRotor("I")

	return Assemble(
		parts.DefaultKeyboard,
		parts.NoPlugboard,
		slow,
		middle,
		fast,
		parts.Reflectors["B"],
		parts.DefaultLightboard)
}

// WithRotors build a new enigma machine, with the specified rotors and reflector.
func WithRotors(slow, middle, fast, reflector string) (Enigma, error) {
	var r1, r2, r3 parts.Rotor
	var err error

	if r1, err = parts.GetRotor(slow); err != nil {
		return nil, err
	}
	if r2, err = parts.GetRotor(middle); err != nil {
		return nil, err
	}
	if r3, err = parts.GetRotor(fast); err != nil {
		return nil, err
	}

	ref, ok := parts.Reflectors[reflector]
	if !ok {
		return nil, errors.New("unknown reflector: '" + reflector + "'")
	}

	return Assemble(
		parts.DefaultKeyboard,
		parts.NoPlugboard,
		r1,
		r2,
		r3,
		ref,
		parts.DefaultLightboard,
	), nil
}

// WithConfig builds a new enigma machine, with default rotors and reflector, using the specified settings.
func WithConfig(ringSetting, windowSetting string) (Enigma, error) {
	e := WithDefaults()
	err := e.Configure(ringSetting, windowSetting)

	return e, err
}

// Assemble builds a new enigma machine, with default config and all parts specified.
// This is the only way to create a machine with a different keyboard, lightboard or plugboard
func Assemble(keyboard parts.Keyboard, plugboard parts.Plugboard, slow, middle, fast parts.Rotor, reflector parts.Reflector, lightboard parts.Lightboard) Enigma {
	enigma := &enigmaImpl{keyboard, plugboard, slow, middle, fast, reflector, lightboard}
	enigma.SetWindow("AAA")
	enigma.SetRing("AAA")

	return Enigma(enigma)
}

func (e *enigmaImpl) Reflector() string {
	return e.reflector.ID()
}

func (e *enigmaImpl) Slow() string {
	return e.slow.ID()
}

func (e *enigmaImpl) Middle() string {
	return e.middle.ID()
}

func (e *enigmaImpl) Fast() string {
	return e.fast.ID()
}

func (e *enigmaImpl) Window() string {
	return string([]rune{
		e.slow.Window(),
		e.middle.Window(),
		e.fast.Window(),
	})
}

func (e *enigmaImpl) SetWindow(settings string) error {
	var runes []rune

	if settings == "" {
		runes = []rune{'A', 'A', 'A'}
	} else {
		runes = []rune(settings)

		if len(runes) != 3 {
			return errors.New("window settings should be 3 characters long (ex: AAA)")
		}

		for _, c := range runes {
			if c < 'A' || c > 'Z' {
				return errors.New("window settings should be specified using only uppercase letters from 'A' to 'Z'")
			}
		}
	}

	e.slow.SetWindow(runes[0])
	e.middle.SetWindow(runes[1])
	e.fast.SetWindow(runes[2])
	return nil
}

func (e *enigmaImpl) Ring() string {
	return string([]rune{
		e.slow.Ring(),
		e.middle.Ring(),
		e.fast.Ring(),
	})
}

func (e *enigmaImpl) SetRing(settings string) error {
	var runes []rune

	if settings == "" {
		runes = []rune{'A', 'A', 'A'}
	} else {
		runes = []rune(settings)

		if len(runes) != 3 {
			return errors.New("ring settings should be 3 characters long (ex: AAA)")
		}

		for _, c := range runes {
			if c < 'A' || c > 'Z' {
				return errors.New("ring settings should be specified using only uppercase letters from 'A' to 'Z'")
			}
		}
	}

	e.slow.SetRing(runes[0])
	e.middle.SetRing(runes[1])
	e.fast.SetRing(runes[2])
	return nil
}

func (e *enigmaImpl) Configure(ringSetting, windowSetting string) error {
	if err := e.SetRing(ringSetting); err != nil {
		return err
	}

	return e.SetWindow(windowSetting)
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
