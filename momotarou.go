package momotarou

func New() *Momotarou {
	return new(Momotarou)
}

type Momotarou struct {
	runner runner
}

type runner interface {
	Run() error
}

func (m *Momotarou) Run() error {
	return m.runner.Run()
}
