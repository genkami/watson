package yaml

import (
	"io"

	"gopkg.in/yaml.v2"

	"github.com/genkami/watson/pkg/types"
)

func Decode(w io.Writer, val *types.Value) error {
	obj := val.ToGoObject()
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	return enc.Encode(obj)
}

func Encode(r io.Reader) (*types.Value, error) {
	var any interface{}
	dec := yaml.NewDecoder(r)
	err := dec.Decode(&any)
	if err != nil {
		return nil, err
	}
	return types.ToValue(any)
}
