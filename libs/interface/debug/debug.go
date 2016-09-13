package debug

type IDebug interface {
	Error(err error)
	Info(format string, a ...interface{})
}
