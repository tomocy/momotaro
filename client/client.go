package client

import (
	"fmt"

	kibidangoPkg "github.com/tomocy/kibidango"
)

func New() *Client {
	return &Client{
		runner: newCLI(),
	}
}

type Client struct {
	runner runner
}

type runner interface {
	Run(args []string) error
}

func (c *Client) Run(args []string) error {
	return c.runner.Run(args)
}

func newFactory(os string) factory {
	if factory, ok := newFor(os).(factory); ok {
		return factory
	}

	return nil
}

type factory interface {
	list() ([]kibidango, error)
	create(spec *kibidangoPkg.Spec) (kibidango, error)
	load(id string) (kibidango, error)
	delete(id string) error
}

func newPrinter(os string) printer {
	if printer, ok := newFor(os).(printer); ok {
		return printer
	}

	return nil
}

type printer interface {
	printAll(kibis []kibidango)
	print(kibi kibidango)
}

type kibidango interface {
	Run(args ...string) error
	Init() error
}

func newFor(os string) interface{} {
	switch os {
	case osLinux:
		return new(linux)
	default:
		return nil
	}
}

const (
	osLinux = "linux"
)

func printHeader() {
	fmt.Println("ID")
}
