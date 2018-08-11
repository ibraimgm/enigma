package parts_test

import (
	"testing"

	"github.com/ibraimgm/enigma/machine/parts"
	"github.com/stretchr/testify/assert"
)

func TestReflectSimple(t *testing.T) {
	var tests = []struct {
		in  int
		out int
	}{
		{6, 19}, // F -> S
		{2, 18}, // B -> R
		{3, 21}, // C -> U
	}

	reflector := parts.Reflectors["B"]

	for _, test := range tests {
		expected := parts.Signal(test.out)
		actual := reflector.Reflect(parts.Signal(test.in))

		assert.Equal(t, expected, actual)
	}
}
