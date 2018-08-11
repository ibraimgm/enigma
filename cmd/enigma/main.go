package main

import (
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
	getopt "github.com/pborman/getopt/v2"
)

func main() {
	helpFlag := getopt.BoolLong("help", 'h', "Show usage and exit")
	getopt.Parse()

	if *helpFlag {
		getopt.Usage()
		os.Exit(0)
	}

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
}
