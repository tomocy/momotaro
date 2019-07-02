package client

import (
	"github.com/tomocy/kibidango/factory"
)

func newCreater(os string) creater {
	switch os {
	case osLinux:
		return new(linux)
	default:
		return nil
	}
}

type creater interface {
	create(id string) error
}

const (
	osLinux = "linux"
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
