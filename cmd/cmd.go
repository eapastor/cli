package cmd

import (
	"github.com/jawher/mow.cli"
	"os"
	"github.com/lastbackend/lb/cmd/command"
	"github.com/lastbackend/lb/libs/io/filesystem"
	"github.com/lastbackend/lb/cmd/config"
	"github.com/lastbackend/lb/cmd/context"
	"github.com/lastbackend/lb/libs/api"
	"github.com/lastbackend/lb/libs/io/debug"
	"github.com/lastbackend/lb/libs/storage"
)

func Execute() {

	app := cli.App("lb", "Last.Backend cli application")

	app.Version("v version", "lb 1.0.0")

	command.InitCommands(app)

	app.Run(os.Args)

}

func init() {

	var cfg = config.Get()
	var ctx = context.Get()

	homedir, err := filesystem.GetHomeDir()
	if err != nil {
		os.Exit(1)
	}

	cfg.LBDir = homedir + "/.lb"

	ctx.API = api.New(cfg.ApiHost)
	ctx.Debug = debug.New(3)
	ctx.Storage.Token = storage.New(cfg.LBDir + "/.tkn")

	// Creating directories and files

	filesystem.MkDir(cfg.LBDir)

}

//func commandNotFound(c *cli.Context, command string) {
//	io.Println (color.Red("Command ") + color.Cyan(command) + color.Red(" not found!"))
//}