// +build !android,!windows

package raylib

import (
	"fmt"
	"os"
)

// Log message types
const (
	LogInfo = iota
	LogWarning
	LogError
	LogDebug
)

var traceDebugMsgs = false

var log_level = -1

func SetLoggingGo(lvl int) {
	log_level = lvl
}

// SetDebug - Set debug messages
func SetDebug(enabled bool) {
	traceDebugMsgs = enabled
}

// TraceLog - Show trace log messages (INFO, WARNING, ERROR, DEBUG)
func TraceLog(msgType int, text string, v ...interface{}) {
	if traceDebugMsgs {
		switch msgType {
			case LogInfo:
				if log_level < LogInfo && log_level > -1 {
					fmt.Printf("INFO: "+text+"\n", v...)
				}
			case LogWarning:
				if log_level < LogWarning && log_level > -1  {
					fmt.Printf("WARNING: "+text+"\n", v...)
				}
			case LogError:
				if log_level < LogError && log_level > -1  {
					fmt.Printf("ERROR: "+text+"\n", v...)
				}
				os.Exit(1)
			case LogDebug:
				if log_level < LogDebug && log_level > -1  {
					fmt.Printf("DEBUG: "+text+"\n", v...)
				}
		}
	}
}

// HomeDir - Returns user home directory
func HomeDir() string {
	return os.Getenv("HOME")
}
