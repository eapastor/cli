package cmd

import (
	"github.com/jawher/mow.cli"
	"os"
	"github.com/lastbackend/cli/cmd/command"
	"github.com/lastbackend/cli/libs/io/filesystem"
	"github.com/lastbackend/cli/cmd/config"
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/api"
	"github.com/lastbackend/cli/libs/io/debug"
	"github.com/lastbackend/cli/libs/storage"
)

func Execute() {

	app := cli.App("lb", "Last.Backend cli application")

	app.Version("v version", "cli 1.0.0")

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