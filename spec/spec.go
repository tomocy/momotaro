package spec

import (
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/tomocy/kibidango"
)

func convertOCISpec(ociSpec *specs.Spec) *spec {
	return &spec{
		process: &kibidango.Process{
			Args: ociSpec.Process.Args,
		},
	}
}

type spec struct {
	process *kibidango.Process
}
