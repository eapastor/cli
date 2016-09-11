package io

import "fmt"

type log struct {}

var logger = log{}

var typeMsg struct {
	info string
	err  string
}

func GetLogger() *log {
	return &logger
}

func (log) Warn(s string) {
	fmt.Println(red(typeMsg.info) + s)
}

func (log) Info(s string) {
	fmt.Println(green("INFO: ") + s)
}