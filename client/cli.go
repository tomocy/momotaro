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
			Name:   "list",
			Action: c.list,
		},
		cliPkg.Command{
			Name:   "create",
			Action: c.create,
		},
		cliPkg.Command{
			Name:   "init",
			Action: c.init,
		},
		cliPkg.Command{
			Name:   "delete",
			Action: c.delete,
		},
	}
}

func (c *cli) Run(args []string) error {
	return c.app.Run(args)
}

func (c *cli) list(ctx *cliPkg.Context) error {
	factory := c.factory()
	kibis, err := factory.list()
	if err != nil {
		return err
	}

	printer := c.printer()
	printer.printAll(kibis)

	return nil
}

func (c *cli) create(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := c.factory()

	kibi, err := factory.create(id)
	if err != nil {
		return err
	}

	return kibi.Run("init", id)
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

func (c *cli) delete(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := c.factory()

	return factory.delete(id)
}

func (c *cli) factory() factory {
	return newFactory(c.os)
}

func (c *cli) printer() printer {
	return newPrinter(c.os)
}
