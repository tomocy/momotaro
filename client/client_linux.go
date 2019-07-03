package client

import (
	factoryPkg "github.com/tomocy/kibidango/factory"
)

type linux struct{}

func (l *linux) create(id string) (kibidango, error) {
	factory := l.factory()
	kibi, err := factory.Manufacture(id)
	if err != nil {
		return nil, err
	}

	if err := factory.Save(kibi); err != nil {
		return nil, err
	}

	return kibi, nil
}

func (l *linux) factory() *factoryPkg.Linux {
	return factoryPkg.ForLinux()
}
