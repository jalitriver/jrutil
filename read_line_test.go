package jrutil

import (
	"bufio"
	"io"
	"slices"
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	data := []struct{
		text string
		expected []string
	}{
		{
			text: "",
			expected: []string{},
		},
		{
			text: "\r",
			expected: []string{"\r"},
		},
		{
			text: "\n",
			expected: []string{"\n"},
		},
		{
			text: "\r\n",
			expected: []string{"\r\n"},
		},
		{
			text: "foo",
			expected: []string{"foo"},
		},
		{
			text: "foo\r",
			expected: []string{"foo\r"},
		},
		{
			text: "foo\n",
			expected: []string{"foo\n"},
		},
		{
			text: "foo\r\n",
			expected: []string{"foo\r\n"},
		},
		{
			text: "foo\rbar",
			expected: []string{"foo\r", "bar"},
		},
		{
			text: "foo\nbar",
			expected: []string{"foo\n", "bar"},
		},
		{
			text: "foo\r\nbar",
			expected: []string{"foo\r\n", "bar"},
		},
		{
			text: "foo\rbar\r",
			expected: []string{"foo\r", "bar\r"},
		},
		{
			text: "foo\nbar\n",
			expected: []string{"foo\n", "bar\n"},
		},
		{
			text: "foo\r\nbar\r\n",
			expected: []string{"foo\r\n", "bar\r\n"},
		},
		{
			text: "foo\rbar\rbaz",
			expected: []string{"foo\r", "bar\r", "baz"},
		},
		{
			text: "foo\nbar\nbaz",
			expected: []string{"foo\n", "bar\n", "baz"},
		},
		{
			text: "foo\r\nbar\r\nbaz",
			expected: []string{"foo\r\n", "bar\r\n", "baz"},
		},
		{
			text: "foo\rbar\rbaz\r",
			expected: []string{"foo\r", "bar\r", "baz\r"},
		},
		{
			text: "foo\nbar\nbaz\n",
			expected: []string{"foo\n", "bar\n", "baz\n"},
		},
		{
			text: "foo\r\nbar\r\nbaz\r\n",
			expected: []string{"foo\r\n", "bar\r\n", "baz\r\n"},
		},
		{
			text: "foo\rbar\nbaz\r\n",
			expected: []string{"foo\r", "bar\n", "baz\r\n"},
		},
	}

	for _, d := range data {
		var actual []string
		r := bufio.NewReader(strings.NewReader(d.text))
		for {
			line, err := ReadLine(r)
			if line != "" {
				actual = append(actual, line)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				t.Errorf("ReadLine: %v", err)
				continue
			}
		}
		if !slices.Equal(actual, d.expected) {
			t.Errorf("ReadLine(%q): expected=%q  actual=%q",
				d.text, d.expected, actual)
		}
	}
}
