package jrutil

import (
	"testing"
)

func TestApplog(t *testing.T) {
	t.Skip()
	err := Applog("app.log", "Hello, %s!", "World")
	if err != nil {
		t.Errorf("TestApplog: %v", err)
	}
}
