package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
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
	save(kibi kibidango) error
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
}

type kibidango interface {
	Run(args ...string) error
	Init() error
	Exec() error
	Spec() *kibidangoPkg.Spec
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

func tableWriter() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "PID", "Command"})

	return table
}

type table struct{}

func (t *table) printAll(kibis []kibidango) {
	writer := tablewriter.NewWriter(os.Stdout)
	writer.SetHeader([]string{"ID", "PID", "Command"})
	for _, kibi := range kibis {
		joined := t.joinSpec(kibi.Spec())
		writer.Append(joined)
	}

	writer.Render()
}

func (t *table) joinSpec(spec *kibidangoPkg.Spec) []string {
	return []string{
		spec.ID,
		fmt.Sprintf("%d", spec.Process.ID),
		strings.Join(spec.Process.Args, " "),
	}
}
