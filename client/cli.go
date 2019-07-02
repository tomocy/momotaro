package client

import (
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

func (c *cli) initCommands() {}

func (c *cli) Run(args []string) error {
	return c.app.Run(args)
}
