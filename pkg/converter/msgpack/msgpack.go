// Package msgpack provides a way to convert MessagePack into types.Value and vice versa.
package msgpack

import (
	"io"

	"github.com/vmihailenco/msgpack"

	"github.com/genkami/watson/pkg/types"
)

func Decode(w io.Writer, val *types.Value) error {
	obj := val.ToGoObject()
	enc := msgpack.NewEncoder(w)
	return enc.Encode(obj)
}

func Encode(r io.Reader) (*types.Value, error) {
	var any interface{}
	dec := msgpack.NewDecoder(r)
	err := dec.Decode(&any)
	if err != nil {
		return nil, err
	}
	return types.ToValue(any)
}
