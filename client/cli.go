package client

import (
	osPkg "os"
	"runtime"

	"github.com/tomocy/kibidango"
	createrPkg "github.com/tomocy/kibidango/creater"
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
	ctner := new(kibidango.Kibidango)

	id := ctx.Args().First()
	if err := ctner.UpdateID(id); err != nil {
		return err
	}

	creater := creater(runtime.GOOS)
	if err := ctner.Create(creater, "init"); err != nil {
		return err
	}

	return save(ctner)
}

func creater(os string) kibidango.Creater {
	switch os {
	case osLinux:
		return createrPkg.ForLinux(osPkg.Stdin, osPkg.Stdout, osPkg.Stderr)
	default:
		return nil
	}
}

func save(ctner *kibidango.Kibidango) error {
	saver := saver(runtime.GOOS)
	return ctner.Save(saver)
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
	ctner := new(kibidango.Kibidango)
	initer := initializer(runtime.GOOS)
	return ctner.Init(initer)
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
