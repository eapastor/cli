package context

import (
	"github.com/lastbackend/cli/libs/interface/debug"
	"github.com/lastbackend/cli/libs/interface/storage"
	"github.com/lastbackend/cli/libs/api"
)

type context struct {
	API     api.API
	Debug   debug.IDebug
	Storage storage.Storage
}

var ctx context

func Get() *context {
	return &ctx
}
