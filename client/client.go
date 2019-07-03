package client

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
	switch os {
	case osLinux:
		return new(linux)
	default:
		return nil
	}
}

type factory interface {
	create(id string) (kibidango, error)
}

type kibidango interface {
	Run(args ...string) error
}

const (
	osLinux = "linux"
)
