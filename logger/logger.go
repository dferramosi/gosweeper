package logger

import (
	"fmt"
	"log"
	"os"
)

var Debug = false

func init() {
	log.SetOutput(os.Stdout)
}

// StartLogger initializes logging
func StartLogger() *os.File {
	cwd, _ := os.Getwd()
	logpath := cwd + "/gosweeper.log"

	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic("Error opening log file")
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return file
}

func DebugLog(msg string) {
	if Debug {
		log.Println(msg)
	}
}

func DebugLogf(msg string, args ...interface{}) {
	DebugLog(fmt.Sprintf(msg, args...))
}

func Log(msg string) {
	log.Println(msg)
}

func Logf(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}
