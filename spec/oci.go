package spec

import (
	"encoding/json"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
)

func ForOCI() *OCI {
	return new(OCI)
}

type OCI struct{}

func (o *OCI) Load(name string) (*Spec, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ociSpec *specs.Spec
	if err := json.NewDecoder(file).Decode(&ociSpec); err != nil {
		return nil, err
	}

	return adaptOCISpec(ociSpec), nil
}

func adaptOCISpec(ociSpec *specs.Spec) *Spec {
	return &Spec{
		Process: &Process{
			Args: ociSpec.Process.Args,
		},
	}
}
