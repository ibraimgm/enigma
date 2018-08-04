package parts_test

import (
	"testing"

	"github.com/ibraimgm/enigma/machine/parts"
	"github.com/stretchr/testify/assert"
)

func TestLightboardUppercase(t *testing.T) {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 1; i <= 26; i++ {
		c := parts.DefaultLightboard.Light(parts.Signal(i))
		assert.Equal(t, int(s[i-1]), int(c))
	}
}

func TestLightboardInvalid(t *testing.T) {
	arr := []int{0, -1, 27}

	for _, v := range arr {
		c := parts.DefaultLightboard.Light(parts.Signal(v))
		assert.Equal(t, 0, int(c))
	}
}
