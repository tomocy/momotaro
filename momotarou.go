package momotarou

func New() *Momotarou {
	return new(Momotarou)
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
