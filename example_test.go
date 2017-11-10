// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package golog_test

import (
	"fmt"
	"bytes"
	. "github.com/mkmueller/golog"
)

func ExampleLogger() {
	var (
		buf    bytes.Buffer
		logger = New(&buf, "logger: ", Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
	// Output:
	// logger: example_test.go:19: Hello, log file!
}

func ExampleLogger_Output() {
	var (
		buf    bytes.Buffer
		logger = New(&buf, "INFO: ", Lshortfile)

		infof = func(info string) {
			logger.Output(2, info)
		}
	)

	infof("Hello world")

	fmt.Print(&buf)
	// Output:
	// INFO: example_test.go:36: Hello world
}


// Print only level 2 and lower messages.
func ExampleLogger_Printf() {

	var buf    bytes.Buffer
	logger := New(&buf, Lshortfile)
	logger.SetLevel(2)

	logger.Printf("%LVL: Zero", 0)
	logger.Printf("%LVL: One", 1)
	logger.Printf("%LVL: Two", 2)
	logger.Printf("%LVL: Three", 3)

	fmt.Print(&buf)
	// Output:
	// example_test.go:51: 0: Zero
	// example_test.go:52: 1: One
	// example_test.go:53: 2: Two
}

// Create a new logger using no arguments. Output to stderr by default.
func ExampleNew() {

	logger := New()
	logger.Printf("Hello stderr")
	// outputs: Hello stderr

	logger = New("/tmp/mylog.log")
	logger.Printf("Hello file")
	// outputs to file: 2017-11-11 01:23:45 Hello file

}

