package jrutil

import (
	"fmt"
	"os"
	"time"
)

// Applog allows you to easily log to a different file to avoid the
// clutter of the system log when debugging.
func Applog(fname string, format string, a ...any) error {
	var err error
	var f *os.File
	var msg string
	var timestamp string

	// Open the log file.
	f, err = os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the log message.
	msg = fmt.Sprintf(format, a...)
	timestamp = time.Now().Format("2006-01-02T15:04:05-07:00")
	_, err = f.WriteString(fmt.Sprintf("%v: %v\n", timestamp, msg))

	return err
}
