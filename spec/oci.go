package spec

import (
	"encoding/json"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/tomocy/kibidango"
)

type OCI struct{}

func (o *OCI) Load(name string) (*kibidango.Spec, error) {
	src, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer src.Close()

	var ociSpec *specs.Spec
	if err := json.NewDecoder(src).Decode(&ociSpec); err != nil {
		return nil, err
	}

	return o.adapt(ociSpec), nil
}

func (o *OCI) adapt(ociSpec *specs.Spec) *kibidango.Spec {
	return &kibidango.Spec{
		Process: &kibidango.Process{
			Args: ociSpec.Process.Args,
		},
	}
}
