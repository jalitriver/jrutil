package jrutil

import (
	"bufio"
	"os"
	"strings"
)

// Prompt prompts on fout for input and then reads a line of text from
// fin.  The line of text is returned with the trailing EOL sequence
// removed.
//
// In the common case of wanting to read and write from and to
// os.Stdin and os.Stdout, they should be wrapped to be buffered as
// follows which should only happen once probably early in your main()
// function:
//
//  fin := bufio.NewReader(os.Stdin)
//  fout := bufio.NewWriter(os.Stdout)
//
// The reason for having to use buffered I/O for stdin and stdout is
// to avoid conflicts with other parts of the code that might need to
// use stdin and stdout.  You cannot mix buffered I/O and non-buffered
// I/O without losing some bytes along the way.  You cannot even
// create multiple buffered I/O wrappers around stdin or stdout
// without losing bytes along the way.
//
// Also see PromptUnbuffered().
func Prompt(
	fin *bufio.Reader,
	fout *bufio.Writer,
	prompt string,
) (string, error) {
	var err error
	var line string

	// Write the prompt.
	_, err = fout.WriteString(prompt)
	if err != nil {
		return "", err
	}
	err = fout.Flush()
	if err != nil {
		return "", err
	}

	// Read line.
	line, err = fin.ReadString('\n')
	if err != nil {
		return line, err
	}
	if len(line) >= 1 {
		line = line[:len(line)-1]
	}

	return line, nil
}

// Prompt prompts on os.Stdout for input and then reads a line of text
// from os.Stdin.  The line of text is returned with the trailing EOL
// sequence removed.
//
// Also see Prompt().
func PromptUnbuffered(prompt string) (string, error) {
	var err error
	var line strings.Builder

	// Write the prompt.
	_, err = os.Stdout.WriteString(prompt)
	if err != nil {
		return "", err
	}

	// Read line.
	ch := make([]byte, 1)
	for {

		// Get the next character.
		_, err = os.Stdin.Read(ch)
		if err != nil {
			return line.String(), err
		}

		// Check for EOL.
		if ch[0] == '\n' {
			break
		}

		// Append this character to the result.
		err = line.WriteByte(ch[0])
		if err != nil {
			return line.String(), err
		}
	}

	return line.String(), err
}
