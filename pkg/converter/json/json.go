package json

import (
	"encoding/json"
	"io"

	"github.com/genkami/watson/pkg/types"
)

func Decode(w io.Writer, val *types.Value) error {
	obj := val.ToGoObject()
	enc := json.NewEncoder(w)
	return enc.Encode(obj)
}

func Encode(r io.Reader) (*types.Value, error) {
	var any interface{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&any)
	if err != nil {
		return nil, err
	}
	return types.ToValue(any)
}
