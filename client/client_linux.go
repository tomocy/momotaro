package client

import (
	factoryPkg "github.com/tomocy/kibidango/factory"
)

type linux struct{}

func (l *linux) create(id string) error {
	factory := factoryPkg.ForLinux()
	kibi, err := factory.Manufacture(id)
	if err != nil {
		return err
	}

	return factory.Save(kibi)
}
