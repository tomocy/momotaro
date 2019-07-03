package client

import (
	"runtime"

	cliPkg "github.com/urfave/cli"
)

func newCLI() *cli {
	c := new(cli)
	c.setUp()
	return c
}

type cli struct {
	os  string
	app *cliPkg.App
}

func (c *cli) setUp() {
	c.os = runtime.GOOS
	c.app = cliPkg.NewApp()
	c.setBasic()
	c.setCommands()
}

func (c *cli) setBasic() {
	c.app.Name = name
	c.app.Usage = usage
	c.app.Version = version
}

const (
	name    = "kibidango"
	usage   = "a client for linux container runtime"
	version = "0.0.1"
)

func (c *cli) setCommands() {
	c.app.Commands = []cliPkg.Command{
		cliPkg.Command{
			Name:   "create",
			Action: c.create,
		},
	}
}

func (c *cli) Run(args []string) error {
	return c.app.Run(args)
}

func (c *cli) create(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := newFactory(c.os)

	return factory.create(id)
}
