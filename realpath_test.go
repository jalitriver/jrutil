package jrutil

import (
	"testing"
)

func TestRealpath(t *testing.T) {
	t.Skip()

	expected := "/usr/lib/modules"

	// Only reasonable on Linux systems.
	actual, err := Realpath("/lib/modules")
	if err != nil {
		t.Errorf("TestRealpath: %v", err)
		return
	}

	// For "merged" Linux systems, /lib/modules now points to /usr/lib/modules.
	if actual != expected {
		t.Errorf(
			"TestRealpath: expected=\"%s\"  actual=\"%s\"",
			expected, actual)
	}
}
