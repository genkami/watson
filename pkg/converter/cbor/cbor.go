package cbor

import (
	"io"

	"github.com/fxamacker/cbor/v2"

	"github.com/genkami/watson/pkg/types"
)

func Decode(w io.Writer, val *types.Value) error {
	obj := val.ToGoObject()
	enc := cbor.NewEncoder(w)
	return enc.Encode(obj)
}

func Encode(r io.Reader) (*types.Value, error) {
	var any interface{}
	dec := cbor.NewDecoder(r)
	err := dec.Decode(&any)
	if err != nil {
		return nil, err
	}
	return types.ToValue(any)
}
