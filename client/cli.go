package client

import (
	"fmt"
	osPkg "os"
	"runtime"

	"github.com/tomocy/kibidango"
	clonerPkg "github.com/tomocy/kibidango/cloner"
	initializerPkg "github.com/tomocy/kibidango/initializer"
	listerPkg "github.com/tomocy/kibidango/lister"
	loaderPkg "github.com/tomocy/kibidango/loader"
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
		{
			Name:   "list",
			Action: list,
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

func list(*cliPkg.Context) error {
	loader := loader(runtime.GOOS)
	lister := lister(runtime.GOOS)
	kibis, err := kibidango.List(lister, loader)
	if err != nil {
		return err
	}

	print(kibis)

	return nil
}

func loader(os string) kibidango.Loader {
	switch os {
	case osLinux:
		return loaderPkg.ForLinux()
	default:
		return nil
	}
}

func lister(os string) kibidango.Lister {
	switch os {
	case osLinux:
		return listerPkg.ForLinux()
	default:
		return nil
	}
}

func print(kibis []*kibidango.Kibidango) {
	printHeader()
	for _, kibi := range kibis {
		printable := printable(*kibi)
		fmt.Println(printable)
	}
}

func printHeader() {
	fmt.Println("ID")
}

type printable kibidango.Kibidango

func (k printable) String() string {
	return fmt.Sprintf("%s", k.ID)
}

const (
	osLinux = "linux"
)
