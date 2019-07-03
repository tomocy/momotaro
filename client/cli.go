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
		cliPkg.Command{
			Name:   "init",
			Action: c.init,
		},
	}
}

func (c *cli) Run(args []string) error {
	return c.app.Run(args)
}

func (c *cli) create(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := c.factory()

	kibi, err := factory.create(id)
	if err != nil {
		return err
	}

	return kibi.Run()
}

func (c *cli) init(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := c.factory()

	kibi, err := factory.load(id)
	if err != nil {
		return err
	}

	return kibi.Init()
}

func (c *cli) factory() factory {
	return newFactory(c.os)
}
