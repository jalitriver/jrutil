package jrutil

import (
	"slices"
	"strings"
	"testing"
)

// TestForEachLineStrict tests ForEachLine() passing in true for the
// "strict" parameter which should work with DOS, Mac, and Unix EOLs.
func TestForEachLineStrict(t *testing.T) {
	data := []struct {
		text     string
		stripEOL bool
		expected []string
	}{
		// Stripped
		{
			text:     "",
			stripEOL: true,
			expected: []string{},
		},
		{
			text:     "\r",
			stripEOL: true,
			expected: []string{""},
		},
		{
			text:     "\n",
			stripEOL: true,
			expected: []string{""},
		},
		{
			text:     "\r\n",
			stripEOL: true,
			expected: []string{""},
		},
		{
			text:     "foo",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\r",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\n",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\r\n",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\rbar",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\nbar",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\r\nbar",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\rbar\r",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\nbar\n",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\r\nbar\r\n",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\rbar\rbaz",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\nbar\nbaz",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\r\nbar\r\nbaz",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\rbar\rbaz\r",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\nbar\nbaz\n",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\r\nbar\r\nbaz\r\n",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\rbar\nbaz\r\n",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},

		// Unstripped
		{
			text:     "",
			stripEOL: false,
			expected: []string{},
		},
		{
			text:     "\r",
			stripEOL: false,
			expected: []string{"\r"},
		},
		{
			text:     "\n",
			stripEOL: false,
			expected: []string{"\n"},
		},
		{
			text:     "\r\n",
			stripEOL: false,
			expected: []string{"\r\n"},
		},
		{
			text:     "foo",
			stripEOL: false,
			expected: []string{"foo"},
		},
		{
			text:     "foo\r",
			stripEOL: false,
			expected: []string{"foo\r"},
		},
		{
			text:     "foo\n",
			stripEOL: false,
			expected: []string{"foo\n"},
		},
		{
			text:     "foo\r\n",
			stripEOL: false,
			expected: []string{"foo\r\n"},
		},
		{
			text:     "foo\rbar",
			stripEOL: false,
			expected: []string{"foo\r", "bar"},
		},
		{
			text:     "foo\nbar",
			stripEOL: false,
			expected: []string{"foo\n", "bar"},
		},
		{
			text:     "foo\r\nbar",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar"},
		},
		{
			text:     "foo\rbar\r",
			stripEOL: false,
			expected: []string{"foo\r", "bar\r"},
		},
		{
			text:     "foo\nbar\n",
			stripEOL: false,
			expected: []string{"foo\n", "bar\n"},
		},
		{
			text:     "foo\r\nbar\r\n",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar\r\n"},
		},
		{
			text:     "foo\rbar\rbaz",
			stripEOL: false,
			expected: []string{"foo\r", "bar\r", "baz"},
		},
		{
			text:     "foo\nbar\nbaz",
			stripEOL: false,
			expected: []string{"foo\n", "bar\n", "baz"},
		},
		{
			text:     "foo\r\nbar\r\nbaz",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar\r\n", "baz"},
		},
		{
			text:     "foo\rbar\rbaz\r",
			stripEOL: false,
			expected: []string{"foo\r", "bar\r", "baz\r"},
		},
		{
			text:     "foo\nbar\nbaz\n",
			stripEOL: false,
			expected: []string{"foo\n", "bar\n", "baz\n"},
		},
		{
			text:     "foo\r\nbar\r\nbaz\r\n",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar\r\n", "baz\r\n"},
		},
		{
			text:     "foo\rbar\nbaz\r\n",
			stripEOL: false,
			expected: []string{"foo\r", "bar\n", "baz\r\n"},
		},
	}

	for _, d := range data {
		var actual []string
		r := strings.NewReader(d.text)
		err := ForEachLine(r, d.stripEOL, /* strict */ true,
			func(line string) (bool, error) {
				actual = append(actual, line)
				return true, nil
			})
		if err != nil {
			t.Errorf("ForEachLine(%q): %v", d.text, err)
			continue
		}
		if !slices.Equal(actual, d.expected) {
			t.Errorf("ForEachLine(%q): expected=%q  actual=%q",
				d.text, d.expected, actual)
		}
	}
}

// TestForEachLineNotStrict tests ForEachLine() passing in false for
// the "strict" parameter which should work DOS and Unix EOLs but not
// Mac EOLs.
func TestForEachLineNotStrict(t *testing.T) {
	data := []struct {
		text     string
		stripEOL bool
		expected []string
	}{
		// Stripped
		{
			text:     "",
			stripEOL: true,
			expected: []string{},
		},
		{
			text:     "\n",
			stripEOL: true,
			expected: []string{""},
		},
		{
			text:     "\r\n",
			stripEOL: true,
			expected: []string{""},
		},
		{
			text:     "foo",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\n",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\r\n",
			stripEOL: true,
			expected: []string{"foo"},
		},
		{
			text:     "foo\nbar",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\r\nbar",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\nbar\n",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\r\nbar\r\n",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},
		{
			text:     "foo\nbar\nbaz",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\r\nbar\r\nbaz",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\nbar\nbaz\n",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\r\nbar\r\nbaz\r\n",
			stripEOL: true,
			expected: []string{"foo", "bar", "baz"},
		},
		{
			text:     "foo\nbar\r\n",
			stripEOL: true,
			expected: []string{"foo", "bar"},
		},

		// Unstripped
		{
			text:     "",
			stripEOL: false,
			expected: []string{},
		},
		{
			text:     "\r",
			stripEOL: false,
			expected: []string{"\r"},
		},
		{
			text:     "\n",
			stripEOL: false,
			expected: []string{"\n"},
		},
		{
			text:     "\r\n",
			stripEOL: false,
			expected: []string{"\r\n"},
		},
		{
			text:     "foo",
			stripEOL: false,
			expected: []string{"foo"},
		},
		{
			text:     "foo\n",
			stripEOL: false,
			expected: []string{"foo\n"},
		},
		{
			text:     "foo\r\n",
			stripEOL: false,
			expected: []string{"foo\r\n"},
		},
		{
			text:     "foo\nbar",
			stripEOL: false,
			expected: []string{"foo\n", "bar"},
		},
		{
			text:     "foo\r\nbar",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar"},
		},
		{
			text:     "foo\nbar\n",
			stripEOL: false,
			expected: []string{"foo\n", "bar\n"},
		},
		{
			text:     "foo\r\nbar\r\n",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar\r\n"},
		},
		{
			text:     "foo\nbar\nbaz",
			stripEOL: false,
			expected: []string{"foo\n", "bar\n", "baz"},
		},
		{
			text:     "foo\r\nbar\r\nbaz",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar\r\n", "baz"},
		},
		{
			text:     "foo\nbar\nbaz\n",
			stripEOL: false,
			expected: []string{"foo\n", "bar\n", "baz\n"},
		},
		{
			text:     "foo\r\nbar\r\nbaz\r\n",
			stripEOL: false,
			expected: []string{"foo\r\n", "bar\r\n", "baz\r\n"},
		},
		{
			text:     "foo\nbar\r\n",
			stripEOL: false,
			expected: []string{"foo\n", "bar\r\n"},
		},
	}

	for _, d := range data {
		var actual []string
		r := strings.NewReader(d.text)
		err := ForEachLine(r, d.stripEOL, /* strict */ false,
			func(line string) (bool, error) {
				actual = append(actual, line)
				return true, nil
			})
		if err != nil {
			t.Errorf("ForEachLine(%q): %v", d.text, err)
			continue
		}
		if !slices.Equal(actual, d.expected) {
			t.Errorf("ForEachLine(%q): expected=%q  actual=%q",
				d.text, d.expected, actual)
		}
	}
}
