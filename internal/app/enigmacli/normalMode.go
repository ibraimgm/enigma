package enigmacli

import (
	"bufio"
	"fmt"
	"io"
)

func runNormalMode(info *parseInfo, stdout io.Writer, stdin io.Reader) error {
	e := info.e

	fmt.Fprintf(stdout, "=>    Rotors: \t%s,%s,%s\n", e.Slow(), e.Middle(), e.Fast())
	fmt.Fprintf(stdout, "=> Reflector: \t%s\n", e.Reflector())
	fmt.Fprintf(stdout, "=>      Ring: \t%s\n", e.Ring())
	fmt.Fprintf(stdout, "=>    Window: \t%s\n", e.Window())
	fmt.Fprintln(stdout, "--- Running in 'normal' mode; EOF to exit ---")

	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		fmt.Fprintln(stdout, e.EncodeMessage(scanner.Text(), info.blockSize))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
