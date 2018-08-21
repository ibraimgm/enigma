package enigmacli

// Run is the main entry point for the command-line enigma interface
func Run() error {
	info, err := parseArgs()
	if err != nil {
		return err
	}

	if info.isHelp {
		return nil
	}

	if info.isInteractive {
		return runInteractiveMode(info)
	}

	return runNormalMode(info)
}
