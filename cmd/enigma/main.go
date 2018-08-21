package main

import (
	"fmt"
	"os"

	"github.com/ibraimgm/enigma/internal/app/enigmacli"
)

func main() {
	if err := enigmacli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "*** Error: %v\n", err)
		os.Exit(1)
	}
}
