package jrutil

import (
	"bufio"
	"io"
)

// ForEachLine invokes the function fn for each line of text in the
// io.Reader r.  This method uses [jrutil.ReadLine()] to get the next
// line so it works with DOS, Mac, and UNIX EOL sequences.  If
// stripEOL is true, the EOLs will be stripped from the line before fn
// is called.  To receive the next line of text, fn must return (true,
// nil).  If fn returns an error, it is forward to the caller as the
// error returned by ForEachLine().
func ForEachLine(
	r io.Reader,
	stripEOL bool,
	fn func(string) (bool, error),
) error {

	// If necessary, wrap file in bufio.Reader to get buffered input.
	br, ok := r.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(r)
	}

	// Iterate over each line of text.  If the last line of text does
	// not have a trailing EOL sequence, both "line" and "err" will be
	// set so you have to handle the remaining characters in "line"
	// before dealing with the error.
	for {
		var more bool
		var fnErr error

		// Read the next line of text.
		line, err := ReadLine(br)

		// Invoke the callback if at least part of the next line of
		// text was read.
		if line != "" {
			if stripEOL {
				more, fnErr = fn(StripEOL(line))
			} else {
				more, fnErr = fn(line)
			}
		}

		// If the callback returned an error, forward it to the caller.
		if fnErr != nil {
			return fnErr
		}

		// Success.
		if !more || err == io.EOF {
			return nil
		}

		// Return an error if ReadString failed.
		if err != nil {
			return err
		}
	}

	return nil
}
