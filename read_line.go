package jrutil

import (
	"bufio"
	"io"
)

// ReadLine returns the next line of text with EOL sequence still
// attached.  Call [jrutil.StripEOL()] to remove the EOL sequence.
// Valid end-of-line sequences are "\r", "\n", or "\r\n".  This method
// assumes the input stream is UTF-8 encoded.  Performance should be
// improved if r is passed in as a bufio.Reader.  Also see
// [ForEachLine()].
func ReadLine(r *bufio.Reader) (string, error) {

	//
	// NOTE #1: Unfortunately, bufio.ReadString('\n') misses lines
	// that end in just '\r' requiring use to manually inspect the
	// input stream one byte at a time.
	//

	//
	// NOTE #2: Even though strings are UTF-8 encoded and we are
	// assuming the input stream is UTF-8 encoded, we can deal in
	// bytes instead of runes because we just need to find the '\r'
	// and '\n' characters which are unambiguously represented in
	// UTF-8 as a byte.
	//

	var ch byte
	var err error
	var result []byte

	for {

		// Get the next byte.
		ch, err = r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		// Check for "\r" or "\r\n" EOLs.
		if ch == '\r' {

			// Add this byte to our result.
			result = append(result, ch)

			// Looking for "\r\n".
			ch, err = r.ReadByte()
			if err != nil {
				if err == io.EOF {
					break
				}
				return "", err
			}

			// If found "\r\n", this line is done.
			if ch == '\n' {
				result = append(result, ch)
				break  // EOL = "\r\n"
			}

			// Did not find "\r\n" so put the character back.
			err = r.UnreadByte()
			if err != nil {
				return "", err
			}

			break  // EOL = "\r"
		}

		// Check for "\n" EOL.
		if ch == '\n' {
			result = append(result, ch)
			break  // EOL = "\n"
		}
		
		// Add this byte to our result.
		result = append(result, ch)
	}

	// Must return "err", not "nil", to indicate io.EOF.
	return string(result), err
}
