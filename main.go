package main

import (
	"flag"

	"gosweeper/gui"

	"gosweeper/logger"
	// _ "gosweeper/settings"
)

var shouldLogFile bool

func init() {
	flag.BoolVar(&shouldLogFile, "log", false, "Enable logging to file (gosweeper.log) in current directory")
	flag.BoolVar(&logger.Debug, "debug", false, "Enable debug mode")
	flag.Parse()
}

func main() {

	if shouldLogFile {
		fileHandler := logger.StartLogger()
		defer fileHandler.Close()
	} else {
		logger.DebugLog("Logging to file disabled")
	}

	gui.Init()
}
