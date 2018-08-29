package enigma_test

import (
	"testing"

	"github.com/ibraimgm/enigma/machine/enigma"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	e := enigma.WithDefaults()
	assert.Equal(t, "AAA", e.Ring())
	assert.Equal(t, "AAA", e.Window())

	e.Configure("", "XXX")
	assert.Equal(t, "AAA", e.Ring())
	assert.Equal(t, "XXX", e.Window())

	e.Configure("YYY", "")
	assert.Equal(t, "YYY", e.Ring())
	assert.Equal(t, "AAA", e.Window())

	e.Configure("YYY", "XXX")
	assert.Equal(t, "YYY", e.Ring())
	assert.Equal(t, "XXX", e.Window())

	e.Configure("", "")
	assert.Equal(t, "AAA", e.Ring())
	assert.Equal(t, "AAA", e.Window())
}

func TestConfigError(t *testing.T) {
	e := enigma.WithDefaults()

	err := e.SetWindow("ZZ")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "window settings should be 3 characters long")

	err = e.SetWindow("012")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "window settings should be specified using only uppercase letters")

	err = e.SetRing("ZZ")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ring settings should be 3 characters long")

	err = e.SetRing("012")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ring settings should be specified using only uppercase letters")

	err = e.Configure("", "ZZ")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "window settings should be 3 characters long")

	err = e.Configure("", "012")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "window settings should be specified using only uppercase letters")

	err = e.Configure("ZZ", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ring settings should be 3 characters long")

	err = e.Configure("012", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ring settings should be specified using only uppercase letters")
}

func TestWithRotorsCreation(t *testing.T) {
	_, err := enigma.WithRotors("XX", "II", "III", "B")
	assert.EqualError(t, err, "unrecognized rotor ID: 'XX'")

	_, err = enigma.WithRotors("I", "XX", "III", "B")
	assert.EqualError(t, err, "unrecognized rotor ID: 'XX'")

	_, err = enigma.WithRotors("I", "II", "XX", "B")
	assert.EqualError(t, err, "unrecognized rotor ID: 'XX'")

	_, err = enigma.WithRotors("I", "II", "III", "XX")
	assert.EqualError(t, err, "unknown reflector: 'XX'")

	e, _ := enigma.WithRotors("I", "II", "III", "C")
	assert.Equal(t, "I", e.Slow())
	assert.Equal(t, "II", e.Middle())
	assert.Equal(t, "III", e.Fast())
	assert.Equal(t, "C", e.Reflector())
}

func testEncodeRunner(t *testing.T, enigma enigma.Enigma, plain, encoded, windowAfter string) {
	expected := []rune(encoded)

	for i, c := range plain {
		result, ok := enigma.Encode(c)
		assert.True(t, ok)
		assert.Equal(t, expected[i], result)
	}

	assert.Equal(t, windowAfter, enigma.Window())
}

func TestWithDefaultsEncode(t *testing.T) {
	testEncodeRunner(t, enigma.WithDefaults(), "WITHDEFAULTS", "BQEYCDNXBGWH", "AAM")
}

func TestWithRotorsEncode(t *testing.T) {
	e, _ := enigma.WithRotors("I", "II", "III", "C")
	testEncodeRunner(t, e, "WITHROTORS", "IFINGPVFKA", "AAK")
}

func TestWithConfigEncode(t *testing.T) {
	e, _ := enigma.WithConfig("RNG", "WND")
	testEncodeRunner(t, e, "WITHCONFIG", "SYAPXFISKX", "WNN")
}

func TestStepping(t *testing.T) {
	var tests = []struct {
		in          rune
		out         rune
		windowAfter string
	}{
		{'A', 'D', "ADQ"},
		{'A', 'Z', "AER"},
		{'A', 'G', "BFS"},
	}

	e := enigma.WithDefaults()
	e.SetWindow("ADP")

	for _, test := range tests {
		actual, ok := e.Encode(test.in)
		assert.True(t, ok)
		assert.Equal(t, test.out, actual)
		assert.Equal(t, test.windowAfter, e.Window())
	}
}

func TestDoubleStep(t *testing.T) {
	e, _ := enigma.WithConfig("AAA", "AEQ")
	testEncodeRunner(t, e, "A", "L", "BFR")
}

func TestEncodeMessage(t *testing.T) {
	var tests = []struct {
		original  string
		blockSize uint
		expected  string
	}{
		{"ENIGMA", 2, "VW JB FI"},
		{"MACHINE", 3, "HTV YSP H"},
		{"SIMULATION", 5, "XQGTT ICZFW"},
	}

	e := enigma.WithDefaults()

	for _, test := range tests {
		e.Configure("", "")
		actual := e.EncodeMessage(test.original, test.blockSize)
		assert.Equal(t, test.expected, actual)
	}
}

func TestEncodeDecode(t *testing.T) {
	e, _ := enigma.WithConfig("SKY", "RIM")
	s := e.EncodeMessage("LZC KR SK", 0)

	e.Configure("SKY", "RIM")
	s = e.EncodeMessage(s, 0)

	// can you guess the message content?
	assert.Equal(t, "LZCKRSK", s)
}
