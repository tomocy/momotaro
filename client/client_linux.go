package client

import (
	"fmt"
	"strings"

	kibidangoPkg "github.com/tomocy/kibidango"
	factoryPkg "github.com/tomocy/kibidango/factory"
)

type linux struct{}

func (l *linux) list() ([]kibidango, error) {
	factory := l.factory()
	kibis, err := factory.List()
	if err != nil {
		return nil, err
	}

	return l.adaptAll(kibis), nil
}

func (l *linux) adaptAll(kibis []*kibidangoPkg.Linux) []kibidango {
	adapteds := make([]kibidango, len(kibis))
	for i, kibi := range kibis {
		adapteds[i] = kibi
	}

	return adapteds
}

func (l *linux) create(spec *kibidangoPkg.Spec) (kibidango, error) {
	factory := l.factory()
	kibi, err := factory.Manufacture(spec)
	if err != nil {
		return nil, err
	}

	if err := factory.Save(kibi); err != nil {
		return nil, err
	}

	return kibi, nil
}

func (l *linux) save(kibi kibidango) error {
	factory := l.factory()
	linux := kibi.(*kibidangoPkg.Linux)

	return factory.Save(linux)
}

func (l *linux) load(id string) (kibidango, error) {
	factory := l.factory()
	return factory.Load(id)
}

func (l *linux) delete(id string) error {
	factory := l.factory()
	return factory.Delete(id)
}

func (l *linux) factory() *factoryPkg.Linux {
	return factoryPkg.ForLinux()
}

func (l *linux) printAll(kibis []kibidango) {
	table := tableWriter()
	for _, kibi := range kibis {
		linux := kibi.(*kibidangoPkg.Linux)
		printable := printableLinux(*linux)
		table.Append(printable.column())
	}

	table.Render()
}

type printableLinux kibidangoPkg.Linux

func (p printableLinux) column() []string {
	spec := p.Spec()
	return []string{
		spec.ID,
		fmt.Sprintf("%d", spec.Process.ID),
		strings.Join(spec.Process.Args, " "),
	}
}
