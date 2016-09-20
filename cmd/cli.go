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
	{
		Name:   "service",
		Usage:  "",
		Before: auth,
		Subcommands: []cli.Command{
			{
				Name:   "deploy",
				Usage:  "Create service by git push",
				Action: routes.Deploy,
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name, n", Usage: "Service name"},
					cli.StringFlag{Name: "region, rg", Usage: "Region for service", Value: "cu"},
					cli.UintFlag{Name: "memory, m", Usage: "Service memory", Value: 128},
				},
				Subcommands: []cli.Command{
					{
						Name:   "git",
						Usage:  "Create service by git repo",
						Action: routes.Deploy,
						Flags: []cli.Flag{
							cli.StringFlag{Name: "name, n", Usage: "Service name"},
							cli.StringFlag{Name: "region, rg", Usage: "Region for service", Value: "cu"},
							cli.UintFlag{Name: "memory, m", Usage: "Service memory", Value: 128},
							cli.StringFlag{Name: "repo, r", Usage: "Git repository"},
							cli.StringFlag{Name: "brunch, b", Usage: "Git brunch", Value: "master"},
						},
					},
					{
						Name:   "docker",
						Usage:  "Create service by docker image",
						Action: routes.Deploy,
						Flags: []cli.Flag{
							cli.StringFlag{Name: "name, n", Usage: "Service name"},
							cli.StringFlag{Name: "region, rg", Usage: "Region for service", Value: "cu"},
							cli.UintFlag{Name: "memory, m", Usage: "Service memory", Value: 128},
							cli.StringFlag{Name: "image, i", Usage: "Docker image"},
						},
					},
					{
						Name:   "jumpstart",
						Usage:  "Create service by teplate",
						Action: routes.Deploy,
						Flags: []cli.Flag{
							cli.StringFlag{Name: "name, n", Usage: "Service name"},
							cli.StringFlag{Name: "region, rg", Usage: "Region for service", Value: "cu"},
							cli.UintFlag{Name: "memory, m", Usage: "Service memory", Value: 128},
							cli.StringFlag{Name: "template, t", Usage: "LB template"},
						},
					},
				},
			},
			{
				Name:   "list",
				Usage:  "Get service list",
				Action: routes.Services,
			},
		},
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
