package client

import (
	osPkg "os"
	"runtime"

	"github.com/tomocy/kibidango"
	clonerPkg "github.com/tomocy/kibidango/cloner"
	initializerPkg "github.com/tomocy/kibidango/initializer"
	saverPkg "github.com/tomocy/kibidango/saver"
	cliPkg "github.com/urfave/cli"
)

func newCLI() *cli {
	c := new(cli)
	c.init()
	return c
}

type cli struct {
	app *cliPkg.App
}

func (c *cli) init() {
	c.app = cliPkg.NewApp()
	c.initBasic()
	c.initCommands()
}

func (c *cli) initBasic() {
	c.app.Name = name
	c.app.Usage = usage
	c.app.Version = version
}

const (
	name    = "kibidango"
	usage   = "a client for linux container runtime"
	version = "0.0.1"
)

func (c *cli) initCommands() {
	c.app.Commands = []cliPkg.Command{
		{
			Name:   "create",
			Action: create,
		},
		{
			Name:   "init",
			Action: initialize,
		},
	}
}

func (c *cli) Run(args []string) error {
	return c.app.Run(args)
}

func create(ctx *cliPkg.Context) error {
	kibi := new(kibidango.Kibidango)

	id := ctx.Args().First()
	if err := kibi.UpdateID(id); err != nil {
		return err
	}

	cloner := cloner(runtime.GOOS)
	if err := kibi.Clone(cloner, "init"); err != nil {
		return err
	}

	return save(kibi)
}

func cloner(os string) kibidango.Cloner {
	switch os {
	case osLinux:
		return clonerPkg.ForLinux(osPkg.Stdin, osPkg.Stdout, osPkg.Stderr)
	default:
		return nil
	}
}

func save(kibi *kibidango.Kibidango) error {
	saver := saver(runtime.GOOS)
	return kibi.Save(saver)
}

func saver(os string) kibidango.Saver {
	switch os {
	case osLinux:
		return saverPkg.ForLinux()
	default:
		return nil
	}
}

func initialize(ctx *cliPkg.Context) error {
	kibi := new(kibidango.Kibidango)
	initer := initializer(runtime.GOOS)
	return kibi.Init(initer)
}

func initializer(os string) kibidango.Initializer {
	switch os {
	case osLinux:
		return initializerPkg.ForLinux("/root/container")
	default:
		return nil
	}
}

const (
	osLinux = "linux"
)
