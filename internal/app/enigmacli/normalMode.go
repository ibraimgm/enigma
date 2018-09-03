package enigmacli

import (
	"bufio"
	"fmt"
	"io"
)

func runNormalMode(info *parseInfo, stdin io.Reader, stdout, file io.Writer) error {
	e := info.e

	if !info.isQuiet {
		fmt.Fprintf(stdout, "=>    Rotors: \t%s,%s,%s\n", e.Slow(), e.Middle(), e.Fast())
		fmt.Fprintf(stdout, "=> Reflector: \t%s\n", e.Reflector())
		fmt.Fprintf(stdout, "=>      Ring: \t%s\n", e.Ring())
		fmt.Fprintf(stdout, "=>    Window: \t%s\n", e.Window())
		fmt.Fprintln(stdout, "--- Running in 'normal' mode; EOF to exit ---")
	}

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		fmt.Fprintln(file, e.EncodeMessage(scanner.Text(), info.blockSize))
	}

	return scanner.Err()
}
