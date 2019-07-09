package client

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

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

func newPrinter(fmt string) printer {
	switch fmt {
	case fmtTable:
		return new(table)
	default:
		return nil
	}
}

const (
	fmtTable = "table"
)

type printer interface {
	printAll(kibis []kibidango)
}

type kibidango interface {
	Run(args ...string) error
	Init() error
	Exec() error
	Kill(sig os.Signal) error
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

func parseSignal(target string) (os.Signal, error) {
	if num, err := strconv.Atoi(target); err == nil {
		return syscall.Signal(num), nil
	}

	trimed := strings.TrimLeft(strings.ToUpper(target), "SIG")
	if signal, ok := signals[trimed]; ok {
		return signal, nil
	}

	return nil, fmt.Errorf("no such signal")
}

var signals = map[string]os.Signal{
	sighup:  syscall.SIGHUP,
	sigint:  syscall.SIGINT,
	sigquit: syscall.SIGQUIT,
	sigill:  syscall.SIGILL,
	sigtrap: syscall.SIGTRAP,
	sigabrt: syscall.SIGABRT,
	sigfpe:  syscall.SIGFPE,
	sigkill: syscall.SIGKILL,
	sigsegv: syscall.SIGSEGV,
	sigpipe: syscall.SIGPIPE,
	sigalrm: syscall.SIGALRM,
	sigterm: syscall.SIGTERM,
}

const (
	sighup  = "HUP"
	sigint  = "INT"
	sigquit = "QUIT"
	sigill  = "ILL"
	sigtrap = "TRAP"
	sigabrt = "ABRT"
	sigfpe  = "FPE"
	sigkill = "KILL"
	sigsegv = "SEGV"
	sigpipe = "PIPE"
	sigalrm = "ALRM"
	sigterm = "TERM"
)
