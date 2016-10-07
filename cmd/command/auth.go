package command

import (
	"github.com/jawher/mow.cli"
	f "github.com/lastbackend/lb/cli"
)

func (c *command) initAuthCommands() {

	var app = c.app

	app.Command("login", "login with lb credentials", func(c *cli.Cmd) {

		c.Action = func() {
			f.Login()
		}

		c.After = func() {
			f.Whoami()
		}

	})

	app.Command("logout", "logout lb", func(c *cli.Cmd) {

		c.Action = func() {
			f.Logout()
		}

	})

	app.Command("whoami", "logout lb", func(c *cli.Cmd) {

		c.Action = func() {
			f.Whoami()
		}

	})

}
