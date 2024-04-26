package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jalitriver/jrutil"
)

func main() {
	var err error
	var name string

	for {	
		name, err = jrutil.PromptUnbuffered("Name: ")
		if name != "" || err == nil {
			name = strings.TrimSpace(name)
			fmt.Printf("%s\n", name)
		}
		if err != nil {
			goto out
		}
	}

out:

	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "*** Error: %v\n", err)
		os.Exit(1)
	}
	
}
