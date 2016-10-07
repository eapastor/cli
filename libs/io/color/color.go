package color

import "github.com/lastbackend/lb/libs/io/filesystem"

func Yellow(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[33m" + s + "\x1b[39m"
	}
	return s
}

func Red(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[31m" + s + "\x1b[39m"
	}
	return s
}

func Green(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[32m" + s + "\x1b[39m"
	}
	return s
}

func Cyan(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[36m" + s + "\x1b[39m"
	}
	return s
}

func White(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[37m" + s + "\x1b[39m"
	}
	return s
}

func Black(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[30m" + s + "\x1b[39m"
	}
	return s
}
