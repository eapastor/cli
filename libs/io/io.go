package io

import (
	"io"
	"os"
	"fmt"
	"github.com/lastbackend/cli/libs/io/color"
)

// Stdout is used to mock stdout for testing
var Stdout io.Writer = os.Stdout

// Stderr is to mock stderr for testing
var Stderr io.Writer = os.Stderr

// InspectOut is used to mock inspect for testing
var InspectOut io.Writer = os.Stderr

// Print is used to replace `fmt.Print()` but can be mocked out for testing.
func Print(a ...interface{}) {
	fmt.Fprint(Stdout, a...)
}

// Printf is used to replace `fmt.Printf()` but can be mocked out for testing.
func Printf(format string, a ...interface{}) {
	fmt.Fprintf(Stdout, format, a...)
}

func Error(message string)  {
	fmt.Fprintf(Stdout, color.Red("ERROR: ") + "%s\n", message)
}

// Println is used to replace `fmt.Println()` but can be mocked out for testing.
func Println(a ...interface{}) {
	fmt.Fprintln(Stdout, a...)
}

func GetString(prompt string) string {
	Print(prompt)
	var s string
	fmt.Scanln(&s)
	return s
}