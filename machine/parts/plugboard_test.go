package parts_test

import (
	"testing"

	"github.com/ibraimgm/enigma/machine/parts"
	"github.com/stretchr/testify/assert"
)

func TestNoPlugboardDoesNotChangeSignal(t *testing.T) {
	for i := 1; i <= 26; i++ {
		expected := parts.Signal(i)
		received := parts.NoPlugboard.Translate(expected)
		assert.Equal(t, expected, received)
	}
}

func TestNoPlugboardIgnoreUnrecognizedSignal(t *testing.T) {
	numbers := []int{-1, 0, 27, 30, 100}

	for _, v := range numbers {
		expected := parts.Signal(v)
		received := parts.NoPlugboard.Translate(expected)
		assert.Equal(t, expected, received)
	}
}

func TestCustomPlugboardSwapsCorrectly(t *testing.T) {
	plugboardTestAux(t, parts.CreatePlugboard("ABCD"))
}

func TestCustomPlugboardIgnoresOdd(t *testing.T) {
	plugboardTestAux(t, parts.CreatePlugboard("ABCDE"))
}

func plugboardTestAux(t *testing.T, board parts.Plugboard) {
	var table = []struct {
		in  int
		out int
	}{
		{1, 2},
		{2, 1},
		{3, 4},
		{4, 3},
		{5, 5},
		{6, 6},
	}

	for _, v := range table {
		input := parts.Signal(v.in)
		expected := parts.Signal(v.out)
		actual := board.Translate(input)
		assert.Equal(t, expected, actual)
	}
}
