package client

import (
	"os"
	"runtime"

	kibidangoPkg "github.com/tomocy/kibidango"
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
			Hidden:    true,
		},
		cliPkg.Command{
			Name:      "start",
			Usage:     "start a command kibidango is waiting to exec",
			ArgsUsage: "id",
			Action:    c.start,
		},
		cliPkg.Command{
			Name:      "kill",
			Usage:     "kill a kibidango of given id with given signal",
			ArgsUsage: "id signal(default: SIGTERM)",
			Action:    c.kill,
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

	printer := c.printer(fmtTable)
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
	if err := kibi.Run("init", spec.ID); err != nil {
		return err
	}

	return factory.save(kibi)
}

func (c *cli) loadSpec(name string) (*kibidangoPkg.Spec, error) {
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

func (c *cli) start(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := c.factory()

	kibi, err := factory.load(id)
	if err != nil {
		return err
	}

	return kibi.Exec()
}

func (c *cli) kill(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	signal, err := c.parseSignal(ctx)
	if err != nil {
		return err
	}
	factory := c.factory()

	kibi, err := factory.load(id)
	if err != nil {
		return err
	}

	return kibi.Kill(signal)
}

func (c *cli) parseSignal(ctx *cliPkg.Context) (os.Signal, error) {
	sigStr := ctx.Args().Get(1)
	if sigStr == "" {
		sigStr = sigterm
	}

	return parseSignal(sigStr)
}

func (c *cli) delete(ctx *cliPkg.Context) error {
	id := ctx.Args().First()
	factory := c.factory()

	return factory.delete(id)
}

func (c *cli) factory() factory {
	return newFactory(c.os)
}

func (c *cli) printer(fmt string) printer {
	return newPrinter(fmt)
}
