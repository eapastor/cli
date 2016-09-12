package cmd

import (
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/cmd/routes"
	"github.com/urfave/cli"
)

var commands []cli.Command = []cli.Command{
	{
		Name:   "login",
		Usage:  "Login with your LastBackend credentials",
		Action: routes.Login,
		After:  routes.Whoami,
	},
	{
		Name:   "logout",
		Usage:  "Logout",
		Action: routes.Logout,
	},
	{
		Name:   "whoami",
		Usage:  "Get username",
		Action: routes.Whoami,
		Before: auth,
	},
}

var flags []cli.Flag = []cli.Flag{

// bool Flags:
}

func auth(c *cli.Context) error {

	ctx := context.Get()

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}

	ctx.Debug.Info("::token %s", token)

	if token != "" {
		return nil
	}

	err = routes.Login(c)
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}

	ctx.Debug.Info("::token %s", token)

	return nil

}
