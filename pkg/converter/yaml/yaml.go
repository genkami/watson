package yaml

import (
	"io"

	"gopkg.in/yaml.v2"

	"github.com/genkami/watson/pkg/converter/any"
	"github.com/genkami/watson/pkg/vm"
)

func Decode(w io.Writer, val *vm.Value) error {
	obj := any.FromValue(val)
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	return enc.Encode(obj)
}
