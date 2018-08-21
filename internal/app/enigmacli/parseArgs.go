package enigmacli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ibraimgm/enigma/machine/enigma"
	"github.com/ibraimgm/enigma/machine/parts"
	getopt "github.com/pborman/getopt/v2"
)

type parseInfo struct {
	e             enigma.Enigma
	isInteractive bool
	isHelp        bool
	blockSize     uint
}

// parseArgs parse command line arguments and returns a new enigma instance and a boolean indicating
// if it should be run in interactive mode and the block size for the coded text
func parseArgs() (*parseInfo, error) {
	helpFlag := getopt.BoolLong("help", 'h', "Show usage and exit")
	interactiveFlag := getopt.BoolLong("interactive", 'i', "Runs in interactive mode (console gui)")
	rotorsOpt := getopt.StringLong("rotors", 'r', "III,II,I", "Comma-separated list of rotors to be used.", "III,II,I")
	reflectorOpt := getopt.StringLong("reflector", 'f', "B", "Reflector to use.", "B")
	ringOpt := getopt.StringLong("ring", 'g', "AAA", "Ring settings to be used.", "ABC")
	windowOpt := getopt.StringLong("window", 'w', "AAA", "Window settings to be used.", "ABC")
	blockOpt := getopt.IntLong("blocksize", 'b', 5, "Block size of the coded text (default: 5)")

	if err := getopt.Getopt(nil); err != nil {
		return nil, err
	}

	if *helpFlag {
		getopt.PrintUsage(os.Stdout)
		fmt.Fprintln(os.Stdout)
		fmt.Fprintln(os.Stdout, "All command-line arguments are optional.")
		fmt.Fprintln(os.Stdout, "By default, enigma run in 'normal' mode, which reads one line from sdtin and outputs encoded text, until EOF is reached.")
		fmt.Fprintln(os.Stdout, "This means that after writing a line and pressing 'Enter', the coded version will be displayed immediately.")
		fmt.Fprintln(os.Stdout, "The coding process will output the characters in 'blocks', whose size can be controlled with the '-b' flag.")
		return &parseInfo{isHelp: true}, nil
	}

	if _, ok := parts.Reflectors[*reflectorOpt]; !ok {
		return nil, errors.New("invalid reflector '" + *reflectorOpt + "'")
	}

	if *blockOpt < 0 {
		return nil, errors.New("blocksize must be equal or greater than zero")
	}

	rotors := strings.Split(*rotorsOpt, ",")
	if len(rotors) != 3 {
		return nil, errors.New("you should specify 3 rotor ID's")
	}

	e, err := enigma.WithRotors(rotors[0], rotors[1], rotors[2], *reflectorOpt)
	if err != nil {
		return nil, err
	}

	err = e.Configure(*ringOpt, *windowOpt)
	if err != nil {
		return nil, err
	}

	return &parseInfo{e, *interactiveFlag, false, uint(*blockOpt)}, nil
}
