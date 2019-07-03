package client

import (
	factoryPkg "github.com/tomocy/kibidango/factory"
)

type linux struct{}

func (l *linux) create(id string) error {
	factory := l.factory()
	kibi, err := factory.Manufacture(id)
	if err != nil {
		return err
	}
	if err := kibi.Run(); err != nil {
		return err
	}

	return factory.Save(kibi)
}

func (l *linux) factory() *factoryPkg.Linux {
	return factoryPkg.ForLinux()
}
