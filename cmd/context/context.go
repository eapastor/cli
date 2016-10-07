package context

import (
	"github.com/lastbackend/lb/libs/interface/debug"
	"github.com/lastbackend/lb/libs/interface/storage"
	"github.com/lastbackend/lb/libs/api"
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
