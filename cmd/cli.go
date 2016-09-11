package cmd

import (
	"github.com/urfave/cli"
	"github.com/lastbackend/cli/cmd/routes"
)

var commands []cli.Command = []cli.Command{
	{
		Name: "login",
		Usage: "Login with your LastBackend credentials",
		Action: routes.Login,
	},
	{
		Name: "logout",
		Usage: "Logout",
		Action: routes.Logout,
	},
	{
		Name:"whoami",
		Usage: "Get username",
		Action:routes.Whoami,
	},
}

var flags []cli.Flag = []cli.Flag{

	// bool Flags:
	cli.BoolFlag{
		Name:   "debug, d",
		Usage:  "Enable debug mode.",
		Hidden: true,
	},

}