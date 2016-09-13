package debug

import (
	"fmt"
	"os"
	"runtime"
	"github.com/lastbackend/cli/libs/io/color"
)

const _NEWLINE = "\n"

type Debug struct {
	Skip int
}

func New(skip int) *Debug {
	return &Debug{Skip: skip}
}

func (d *Debug) Error(err error) {

	d.printf(color.Red(err.Error()))

}

func (d *Debug) Info(format string, a ...interface{}) {

	d.printf(color.Yellow(format), a...)

}

func (d *Debug) printf(format string, a ...interface{}) {

	if isDebug() {
		fmt.Printf("%s: ", color.Cyan(fileInfo(d.Skip)))
		fmt.Printf(format + _NEWLINE, a...)
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
