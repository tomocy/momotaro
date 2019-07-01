package spec

import (
	"encoding/json"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/tomocy/kibidango"
)

func ForOCI(kibi *kibidango.Kibidango) *OCI {
	return &OCI{
		kibi: kibi,
	}
}

type OCI struct {
	kibi *kibidango.Kibidango
}

func (o *OCI) Load(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	var ociSpec *specs.Spec
	if err := json.NewDecoder(file).Decode(&ociSpec); err != nil {
		return err
	}

	return o.adaptSpec(ociSpec)
}

func (o *OCI) adaptSpec(ociSpec *specs.Spec) error {
	converted := convertOCISpec(ociSpec)
	return adaptSpec(o.kibi, converted)
}

func adaptSpec(kibi *kibidango.Kibidango, spec *spec) error {
	if err := kibi.UpdateProcess(spec.process); err != nil {
		return err
	}

	return nil
}
