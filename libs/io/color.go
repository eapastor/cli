package io

import "github.com/lastbackend/cli/libs/io/filesystem"

func yellow(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[33m" + s + "\x1b[39m"
	}
	return s
}

func red(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[31m" + s + "\x1b[39m"
	}
	return s
}

func green(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[32m" + s + "\x1b[39m"
	}
	return s
}

func cyan(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[36m" + s + "\x1b[39m"
	}
	return s
}

func white(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[37m" + s + "\x1b[39m"
	}
	return s
}

func black(s string) string {
	if !filesystem.IsWindows() {
		return "\x1b[30m" + s + "\x1b[39m"
	}
	return s
}
