package jrutil

// StripEOL returns the line of text with its EOL sequence removed.
// This function strips Unix ('\n'), DOS ('\r\n'), and Mac ('\r') EOL
// sequences.
func StripEOL(line string) string {
	// Even though we are dealing with UTF-8 strings, '\r' and '\n'
	// are both represented as a single byte in UTF-8 so there is no
	// need to convert to runes.
	count := len(line)
	if count >= 2 {
		if line[count-2] == '\r' && line[count-1] == '\n' {
			return line[:count-2]
		}
	}
	if count >= 1 {
		if line[count-1] == '\r' || line[count-1] == '\n' {
			return line[:count-1]
		}
	}
	return line
}
