package main

func main() {
	e, isInteractive, blockSize := parseArgs()

	if isInteractive {
		runInteractiveMode(e)
	} else {
		runNormalMode(e, blockSize)
	}
}
