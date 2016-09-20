package cmd

import (
	"github.com/lastbackend/cli/cmd/config"
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/io/filesystem"
	"github.com/urfave/cli"
	"os"
	"github.com/lastbackend/cli/libs/storage"
	"github.com/lastbackend/cli/libs/api"
	"github.com/lastbackend/cli/libs/io/debug"
	"github.com/lastbackend/cli/libs/io"
	"github.com/lastbackend/cli/libs/io/color"
)

func Start() {

	config.Get()

	initialize()

	app := cli.NewApp()
	app.Name = "lb"
	app.Version = "dev"
	app.Usage = "last.backend cli application"

	app.CommandNotFound = commandNotFound

	app.Flags = flags
	app.Commands = commands

	app.Run(os.Args)

}

func initialize() {

	var cfg = config.Get()
	var ctx = context.Get()

	homedir, err := filesystem.GetHomeDir()
	if err != nil {
		os.Exit(1)
	}

	cfg.LBDir = homedir + "/.LB"

	ctx.API = api.New(cfg.ApiHost)
	ctx.Debug = debug.New(3)
	ctx.Storage.Token = storage.New(cfg.LBDir + "/.tkn")

	// Creating directories and files

	filesystem.MkDir(cfg.LBDir)

}

func commandNotFound(c *cli.Context, command string) {
	io.Println (color.Red("Command ") + color.Cyan(command) + color.Red(" not found!"))
}