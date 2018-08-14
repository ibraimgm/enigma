package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ibraimgm/enigma/machine/enigma"
	"github.com/ibraimgm/enigma/machine/parts"
	getopt "github.com/pborman/getopt/v2"
)

// parseArgs parse command line arguments and returns a new enigma instance and a boolean indicating
// if it should be run in interactive mode and the block size for the coded text
func parseArgs() (enigma.Enigma, bool, uint) {
	helpFlag := getopt.BoolLong("help", 'h', "Show usage and exit")
	interactiveFlag := getopt.BoolLong("interactive", 'i', "Runs in interactive mode (console gui)")
	rotorsOpt := getopt.StringLong("rotors", 'r', "III,II,I", "Comma-separated list of rotors to be used.", "III,II,I")
	reflectorOpt := getopt.StringLong("reflector", 'f', "B", "Reflector to use.", "B")
	ringOpt := getopt.StringLong("ring", 'g', "AAA", "Ring settings to be used.", "ABC")
	windowOpt := getopt.StringLong("window", 'w', "AAA", "Window settings to be used.", "ABC")
	blockOpt := getopt.IntLong("blocksize", 'b', 5, "Block size of the coded text (default: 5)")
	getopt.Parse()

	if *helpFlag {
		getopt.PrintUsage(os.Stdout)
		fmt.Fprintln(os.Stdout)
		fmt.Fprintln(os.Stdout, "All command-line arguments are optional.")
		fmt.Fprintln(os.Stdout, "By default, enigma run in 'normal' mode, which reads one line from sdtin and outputs encoded text, until EOF is reached.")
		fmt.Fprintln(os.Stdout, "This means that after writing a line and pressing 'Enter', the coded version will be displayed immediately.")
		fmt.Fprintln(os.Stdout, "The coding process will output the characters in 'blocks', whose size can be controlled with the '-b' flag.")
		os.Exit(0)
	}

	if _, ok := parts.Reflectors[*reflectorOpt]; !ok {
		fmt.Fprintf(os.Stderr, "*** Error: Invalid reflector '%s'.\n", *reflectorOpt)
		os.Exit(1)
	}

	if *blockOpt < 0 {
		fmt.Fprintf(os.Stderr, "*** Error: blocksize must be equal or greater than zero.")
		os.Exit(1)
	}

	rotors := strings.Split(*rotorsOpt, ",")
	if len(rotors) != 3 {
		fmt.Fprintln(os.Stderr, "*** Error: You should specify 3 rotor ID's.")
		os.Exit(1)
	}

	e, err := enigma.WithRotors(rotors[0], rotors[1], rotors[2], *reflectorOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** Error: %v.\n", err)
		os.Exit(1)
	}

	err = e.Configure(*ringOpt, *windowOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "*** Error: %v.\n", err)
		os.Exit(1)
	}

	return e, *interactiveFlag, uint(*blockOpt)
}
