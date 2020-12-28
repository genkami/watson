package yaml

import (
	"errors"
	"io"

	"gopkg.in/yaml.v2"

	"github.com/genkami/watson/pkg/types"
)

func Decode(w io.Writer, val *types.Value) error {
	obj := val.ToGoObject()
	enc := yaml.NewEncoder(w)
	defer enc.Close()
	arr, ok := obj.([]interface{})
	if !ok {
		return enc.Encode(obj)
	}
	for _, v := range arr {
		err := enc.Encode(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Encode(r io.Reader) (*types.Value, error) {
	var any interface{}
	dec := yaml.NewDecoder(r)
	results := make([]*types.Value, 0)
	for {
		err := dec.Decode(&any)
		if err != nil {
			if errors.Is(err, io.EOF) && len(results) > 0 {
				break
			}
			return nil, err
		}
		v, err := types.ToValue(any)
		if err != nil {
			return nil, err
		}
		results = append(results, v)
	}
	if len(results) == 1 {
		return results[0], nil
	} else {
		return types.NewArrayValue(results), nil
	}
}
