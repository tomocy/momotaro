package client

import (
	"github.com/tomocy/kibidango/factory"
)

type linux struct{}

func (l *linux) create(id string) error {
	factory := factory.ForLinux()
	kibi, err := factory.Manufacture(id)
	if err != nil {
		return err
	}

	return factory.Save(kibi)
}
