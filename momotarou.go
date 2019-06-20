package momotarou

import "github.com/tomocy/momotarou/client"

func New() *Momotarou {
	return &Momotarou{
		runner: client.NewCLI(),
	}
}

type Momotarou struct {
	runner runner
}

type runner interface {
	Run(args []string) error
}

func (m *Momotarou) Run(args []string) error {
	return m.runner.Run(args)
}
