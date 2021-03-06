package command

import (
	"github.com/jawher/mow.cli"
	f "github.com/lastbackend/cli/cli"
)

func (c *command) initServiceCommands() {

	var app = c.app

	app.Command("service", "", func(c *cli.Cmd) {

		c.Before = func() {
			f.Auth()
		}

		c.Command("deploy", "Create service by git push", func(c *cli.Cmd) {

			c.Spec = "NAME [-r][-m][-j | -g | -i]"

			name := c.StringArg("NAME", "", "Service name")
			region := c.StringOpt("r region", "cu", "Region for your service")
			memory := c.IntOpt("m memory", 128, "Service memory")
			template := c.StringOpt("j jumpstart", "", "Create service by teplate")
			git := c.StringOpt("g gitrepo", "", "Create service by git repo")
			docker := c.StringOpt("i image", "", "Create service by docker image")

			uintMemory := uint(*memory)

			c.Action = func() {
				f.SERVICE.Deploy(*name, *region, uintMemory, *template, *git, *docker)
			}

			/*			Subcommands: []cli.Command{
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

						c.Command("git", "Create service by git repo", func(c *cli.Cmd) {

							c.Action = func() {
								fmt.Printf("[GIT] Service deploy with name = %s, region = %s and memory = %d", *name, *region, *memory)
							}

						})
						c.Command("docker", "Create service by docker image", func(c *cli.Cmd) {

							c.Action = func() {
								fmt.Printf("[Docker] Service deploy with name = %s, region = %s and memory = %d", *name, *region, *memory)
							}

						})
						c.Command("jumpstart", "Create service by teplate", func(c *cli.Cmd) {

							c.Action = func() {
								fmt.Printf("[Jumpstart] Service deploy with name = %s, region = %s and memory = %d", *name, *region, *memory)
							}

						})*/

		})

		c.Command("list", "Dislay service list", func(c *cli.Cmd) {

			c.Action = func() {
				f.SERVICE.List()
			}

		})

		c.Command("start", "Start service", func(c *cli.Cmd) {

			c.Spec = "NAME"

			name := c.StringArg("NAME", "", "Service name")

			c.After = func() {
				f.SERVICE.Get(*name)
			}

			c.Action = func() {
				f.SERVICE.Start(*name)
			}

		})

		c.Command("stop", "Stop service", func(c *cli.Cmd) {

			c.Spec = "NAME"

			name := c.StringArg("NAME", "", "Service name")

			c.After = func() {
				f.SERVICE.Get(*name)
			}

			c.Action = func() {
				f.SERVICE.Stop(*name)
			}

		})

		c.Command("restart", "Restart service", func(c *cli.Cmd) {

			c.Spec = "NAME"

			name := c.StringArg("NAME", "", "Service name")

			c.After = func() {
				f.SERVICE.Get(*name)
			}

			c.Action = func() {
				f.SERVICE.Restart(*name)
			}

		})

		c.Command("remove", "Remove service", func(c *cli.Cmd) {

			c.Spec = "NAME"

			name := c.StringArg("NAME", "", "Service name")

			c.After = func() {
				f.SERVICE.List()
			}

			c.Action = func() {
				f.SERVICE.Remove(*name)
			}

		})

		c.Command("get", "Display service by name", func(c *cli.Cmd) {

			c.Spec = "NAME"

			name := c.StringArg("NAME", "", "Service name")

			c.Action = func() {
				f.SERVICE.Get(*name)
			}

		})

	})

}
