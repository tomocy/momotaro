package client

import (
	"runtime"

	"github.com/tomocy/momotaro/spec"
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
			Usage:  "list all kibidangos",
			Action: c.list,
		},
		cliPkg.Command{
			Name:      "create",
			Usage:     "create a kibidango with give id",
			ArgsUsage: "id",
			Action:    c.create,
		},
		cliPkg.Command{
			Name:      "init",
			Usage:     "initialize a kibidango with given id",
			ArgsUsage: "id",
			Action:    c.init,
		},
		cliPkg.Command{
			Name:      "delete",
			Usage:     "delete a kibidango with given id",
			ArgsUsage: "id",
			Action:    c.delete,
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
	factory := c.factory()
	spec, err := c.loadSpec("./config.json")
	if err != nil {
		return err
	}
	spec.ID = ctx.Args().First()

	kibi, err := factory.create(spec)
	if err != nil {
		return err
	}

	return kibi.Run("init", spec.ID)
}

func (c *cli) loadSpec(name string) (*spec.Spec, error) {
	loader := new(spec.OCI)
	return loader.Load(name)
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
