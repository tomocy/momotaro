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
