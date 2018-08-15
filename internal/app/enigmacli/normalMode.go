package enigmacli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ibraimgm/enigma/machine/enigma"
)

func runNormalMode(e enigma.Enigma, blockSize uint) {
	fmt.Fprintf(os.Stdout, "=> Rotors: \t%s,%s,%s\n", e.Slow(), e.Middle(), e.Fast())
	fmt.Fprintf(os.Stdout, "=>   Ring: \t%s\n", e.Ring())
	fmt.Fprintf(os.Stdout, "=> Window: \t%s\n", e.Window())
	fmt.Fprintln(os.Stdout, "--- Running in 'normal' mode; EOF to exit ---")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(e.EncodeMessage(scanner.Text(), blockSize))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "*** Error: %v", err)
	}
}
