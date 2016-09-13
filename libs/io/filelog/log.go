package filelog

import "fmt"

type log struct {}

var logger = log{}

func GetLogger() *log {
	return &logger
}