package client

import (
	osPkg "os"
	"runtime"

	"github.com/tomocy/kibidango/engine/container"
	createrPkg "github.com/tomocy/kibidango/engine/creater"
	initializerPkg "github.com/tomocy/kibidango/engine/initializer"
	"github.com/urfave/cli"
)

func NewCLI() *CLI {
	c := new(CLI)
	c.init()
	return c
}

type CLI struct {
	app *cli.App
}

func (c *CLI) init() {
	c.app = cli.NewApp()
	c.initBasic()
	c.initCommands()
}

func (c *CLI) initBasic() {
	c.app.Name = name
	c.app.Usage = usage
	c.app.Version = version
}

const (
	name    = "kibidango"
	usage   = "a client for linux container runtime"
	version = "0.0.1"
)

func (c *CLI) initCommands() {
	c.app.Commands = []cli.Command{
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

func (c *CLI) Run(args []string) error {
	return c.app.Run(args)
}

func create(ctx *cli.Context) error {
	ctner := new(container.Container)
	creater := creater(runtime.GOOS)
	return ctner.Create(creater, "init")
}

func creater(os string) container.Creater {
	switch os {
	case osLinux:
		return createrPkg.ForLinux(osPkg.Stdin, osPkg.Stdout, osPkg.Stderr)
	default:
		return nil
	}
}

func initialize(ctx *cli.Context) error {
	ctner := new(container.Container)
	initer := initializer(runtime.GOOS)
	return ctner.Init(initer)
}

func initializer(os string) container.Initializer {
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
