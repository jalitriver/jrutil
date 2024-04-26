package jrutil

import (
	"testing"
)

func TestStripEOL(t *testing.T) {
	data := []struct {
		line     string
		expected string
	}{
		{line: "", expected: ""},
		{line: "\r", expected: ""},
		{line: "\n", expected: ""},
		{line: "\r\n", expected: ""},
		{line: "a", expected: "a"},
		{line: "a\r", expected: "a"},
		{line: "a\n", expected: "a"},
		{line: "a\r\n", expected: "a"},
		{line: "foo", expected: "foo"},
		{line: "foo\r", expected: "foo"},
		{line: "foo\n", expected: "foo"},
		{line: "foo\r\n", expected: "foo"},
	}

	for _, d := range data {
		actual := StripEOL(d.line)
		if actual != d.expected {
			t.Errorf("StripEOL(%q): expected=%q  actual=%q",
				d.line, d.expected, actual)
		}
	}
}
