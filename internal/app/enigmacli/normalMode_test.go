package enigmacli

import (
	"errors"
	"strings"
	"testing"

	"github.com/ibraimgm/enigma/machine/enigma"
	"github.com/stretchr/testify/assert"
)

func TestNormalModeOK(t *testing.T) {
	stdin := strings.NewReader("enigma")
	stdout := &strings.Builder{}
	info := &parseInfo{
		e:         enigma.WithDefaults(),
		blockSize: 5,
	}

	err := runNormalMode(info, stdin, stdout, stdout)
	assert.NoError(t, err)

	output := stdout.String()
	assert.Contains(t, output, "--- Running in 'normal' mode; EOF to exit ---")
	assert.Contains(t, output, "VWJBF I")
}

type mockReader struct{}

func (r *mockReader) Read(p []byte) (int, error) {
	return 0, errors.New("some I/O error happened")
}

func TestNormalModeError(t *testing.T) {
	stdin := &mockReader{}
	stdout := &strings.Builder{}
	info := &parseInfo{
		e:         enigma.WithDefaults(),
		blockSize: 5,
	}

	err := runNormalMode(info, stdin, stdout, stdout)
	assert.EqualError(t, err, "some I/O error happened")

	output := stdout.String()
	assert.Contains(t, output, "--- Running in 'normal' mode; EOF to exit ---")
}
