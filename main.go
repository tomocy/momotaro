package main

import (
	"fmt"
	"os"

	"github.com/tomocy/momotarou/client"
)

func main() {
	client := client.New()
	if err := client.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run: %s\n", err)
		os.Exit(1)
	}
}
