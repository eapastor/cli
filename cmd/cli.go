package cmd

/*
import (
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/cmd/cli"
	"github.com/urfave/cli"
)

var commands []cli.Command = []cli.Command{
	{
		Name:   "login",
		Usage:  "Login with your LastBackend credentials",
		Action: cli.Login,
		After:  cli.Whoami,
	},
	{
		Name:   "logout",
		Usage:  "Logout",
		Action: cli.Logout,
	},
	{
		Name:   "whoami",
		Usage:  "Get username",
		Action: cli.Whoami,
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
				Action: cli.Deploy,
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name, n", Usage: "Service name"},
					cli.StringFlag{Name: "region, rg", Usage: "Region for service", Value: "cu"},
					cli.UintFlag{Name: "memory, m", Usage: "Service memory", Value: 128},
				},
				Subcommands: []cli.Command{
					{
						Name:   "git",
						Usage:  "Create service by git repo",
						Action: cli.Deploy,
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
						Action: cli.Deploy,
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
						Action: cli.Deploy,
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
				Action: cli.Services,
			},
		},
	},
}

var flags []cli.Flag = []cli.Flag{

// bool Flags:
}


*/
