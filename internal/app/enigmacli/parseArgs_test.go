package enigmacli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseArgsError(t *testing.T) {
	tests := []struct {
		args    []string
		message string
	}{
		{[]string{"cmd", "-b"}, "missing parameter for -b"},
		{[]string{"cmd", "-b", "-1"}, "blocksize must be equal or greater than zero"},
		{[]string{"cmd", "-b", "X"}, "not a valid number: X"},
		{[]string{"cmd", "-f", "X"}, "invalid reflector 'X'"},
		{[]string{"cmd", "-r", "I,II"}, "you should specify 3 rotor ID's"},
		{[]string{"cmd", "-r", "X,Y,Z"}, "unrecognized rotor ID: 'X'"},
		{[]string{"cmd", "-g", "XY"}, "ring settings should be 3 characters long (ex: AAA)"},
		{[]string{"cmd", "-g", "0YZ"}, "ring settings should be specified using only uppercase letters from 'A' to 'Z'"},
		{[]string{"cmd", "-w", "XY"}, "window settings should be 3 characters long (ex: AAA)"},
		{[]string{"cmd", "-w", "0YZ"}, "window settings should be specified using only uppercase letters from 'A' to 'Z'"},
	}

	for _, test := range tests {
		_, err := parseArgs(test.args, nil)
		assert.EqualError(t, err, test.message)
	}
}

func TestParseArgsHelp(t *testing.T) {
	stdout := &strings.Builder{}
	info, err := parseArgs([]string{"cmd", "-h"}, stdout)
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.True(t, info.isHelp)

	output := stdout.String()
	assert.Contains(t, output, "All command-line arguments are optional.")
}

func TestParseArgsOK(t *testing.T) {
	info, err := parseArgs([]string{"cmd", "-b", "3", "-i"}, nil)
	assert.NoError(t, err)
	assert.NotNil(t, info)

	assert.NotNil(t, info.e)
	assert.False(t, info.isHelp)
	assert.True(t, info.isInteractive)
	assert.Equal(t, info.blockSize, uint(3))
}
