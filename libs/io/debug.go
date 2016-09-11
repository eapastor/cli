package io

import (
	"fmt"
	"os"
	"runtime"
)

type Debug struct {
	Skip int
}

func New(skip int) *Debug {
	return &Debug{Skip: skip}
}

func (d *Debug) Error(err error) {

	d.printf(red(err.Error()))

}

func (d *Debug) Info(format string, a ...interface{}) {

	d.printf(black(format), a...)

}

func (d *Debug) printf(format string, a ...interface{}) {

	if isDebug() {
		fmt.Printf("%s: ", cyan(fileInfo(d.Skip)))
		fmt.Printf(format, a...)
	}

}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func isDebug() bool {

	if os.Getenv("LB_DEBUG_MODE") == "1" {
		return true
	}

	return false

}
