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
