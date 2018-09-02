package enigmacli

import (
	"bufio"
	"os"
)

// Run is the main entry point for the command-line enigma interface
func Run() error {
	info, err := parseArgs(os.Args, os.Stdout)
	if err != nil {
		return err
	}

	if info.isHelp {
		return nil
	}

	outputFile := os.Stdout

	if info.fileName != "" {
		outputFile, err = os.Create(info.fileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		outputFile := bufio.NewWriter(outputFile)
		defer outputFile.Flush()
	}

	return runNormalMode(info, os.Stdin, os.Stdout, outputFile)
}
