package context

import (
	"github.com/lastbackend/cli/libs/interface/debug"
	"github.com/lastbackend/cli/libs/interface/storage"
	"github.com/lastbackend/cli/libs/api"
)

type Context struct {
	API     api.API
	Debug   debug.IDebug
	Storage storage.Storage
}

var ctx Context

func Get() *Context {
	return &ctx
}
