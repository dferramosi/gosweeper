package logger

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartLogger(t *testing.T) {
	file := StartLogger()
	defer file.Close()

	assert.NotNil(t, file)
	_, err := os.Stat(file.Name())
	assert.NoError(t, err)
}

func TestDebugLog(t *testing.T) {
	Debug = true
	defer func() { Debug = false }()

	// Capture the original output
	originalOutput := log.Writer()
	defer log.SetOutput(originalOutput)

	// Capture the output of the log
	r, w, _ := os.Pipe()
	log.SetOutput(w)

	DebugLog("test debug log")

	w.Close()
	out, _ := io.ReadAll(r)

	assert.Contains(t, string(out), "test debug log")
}

func TestDebugLogf(t *testing.T) {
	Debug = true
	defer func() { Debug = false }()

	// Capture the original output
	originalOutput := log.Writer()
	defer log.SetOutput(originalOutput)

	// Capture the output of the log
	r, w, _ := os.Pipe()
	log.SetOutput(w)

	DebugLogf("test debug log %d", 123)

	w.Close()
	out, _ := io.ReadAll(r)

	assert.Contains(t, string(out), "test debug log 123")
}

func TestLog(t *testing.T) {
	// Capture the original output
	originalOutput := log.Writer()
	defer log.SetOutput(originalOutput)

	// Capture the output of the log
	r, w, _ := os.Pipe()
	log.SetOutput(w)

	Log("test log")

	w.Close()
	out, _ := io.ReadAll(r)

	assert.Contains(t, string(out), "test log")
}

func TestLogf(t *testing.T) {
	// Capture the original output
	originalOutput := log.Writer()
	defer log.SetOutput(originalOutput)

	// Capture the output of the log
	r, w, _ := os.Pipe()
	log.SetOutput(w)

	Logf("test log %d", 123)

	w.Close()
	out, _ := io.ReadAll(r)

	assert.Contains(t, string(out), "test log 123")
}
