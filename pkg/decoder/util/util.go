// Package util provides utility functions for converting vm.Value into various formats.
package util

import (
	"fmt"

	"github.com/genkami/watson/pkg/vm"
)

// ToObject converts vm.Value into one of the following type:
// * int64
// * string
// * bool
// * (interface{})(nil)
// * map[string]interface{} (the value of which is also one or many of these types)
func ToObject(val *vm.Value) interface{} {
	switch val.Kind {
	case vm.KInt:
		return val.Int
	case vm.KString:
		return string(val.String)
	case vm.KObject:
		obj := map[string]interface{}{}
		for k, v := range val.Object {
			obj[k] = ToObject(v)
		}
		return obj
	case vm.KBool:
		return val.Bool
	case vm.KNil:
		return nil
	default:
		panic(fmt.Errorf("invalid kind: %d", val.Kind))
	}
}
