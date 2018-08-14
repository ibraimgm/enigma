package main

import (
	"time"

	"github.com/ibraimgm/enigma/machine/enigma"
	termbox "github.com/nsf/termbox-go"
)

func runInteractiveMode(e enigma.Enigma) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)

	for i, c := range "Hello world!" {
		termbox.SetCell(i, 0, c, termbox.ColorWhite, termbox.ColorDefault)
	}

	termbox.Flush()
	time.Sleep(1 * time.Second)
	panic("Interactive mode not implemented yet!")
}
