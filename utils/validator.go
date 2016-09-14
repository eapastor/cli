package utils

import "regexp"

func IsServiceName(s string) bool {

	reg, _ := regexp.Compile("[a-z0-9]+(?:[_-][a-z0-9]+)*")
	str := reg.FindStringSubmatch(s)
	if len(str) == 1 && str[0] == s && len(s) >= 4 && len(s) <= 64 {
		return true
	}
	return false

}

const _DIV_MEMORY uint = 128

func IsMemory(num uint) bool {

	if num % _DIV_MEMORY == 0 {
		return true
	}
	return false

}