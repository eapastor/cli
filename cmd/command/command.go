package command

import (
	"github.com/jawher/mow.cli"
)

type command struct {
	app *cli.Cli
}

var cmd = &command{}

func InitCommands(app *cli.Cli)  {

	cmd.app = app
	cmd.initAuthCommands()
	cmd.initServiceCommands()

}
