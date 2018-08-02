package parts_test

import (
	"testing"

	"github.com/ibraimgm/enigma/machine/parts"
	"github.com/stretchr/testify/assert"
)

func TestDefaultKeyboardUppercase(t *testing.T) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, c := range str {
		s, ok := parts.DefaultKeyboard.InputKey(c)

		assert.True(t, ok)
		assert.Equal(t, i+1, int(s))
	}
}

func TestDefaultKeyboardLowercase(t *testing.T) {
	str := "abcdefghijklmnopqrstuvwxyz"

	for i, c := range str {
		s, ok := parts.DefaultKeyboard.InputKey(c)

		assert.True(t, ok)
		assert.Equal(t, i+1, int(s))
	}
}

func TestDefaultKeyboardNoNumbers(t *testing.T) {
	str := "0123456789"

	for _, c := range str {
		_, ok := parts.DefaultKeyboard.InputKey(c)

		assert.False(t, ok)
	}
}

func TestDefaultKeyboardNoCommonChars(t *testing.T) {
	str := " ./?:!@,"

	for _, c := range str {
		_, ok := parts.DefaultKeyboard.InputKey(c)

		assert.False(t, ok)
	}
}
