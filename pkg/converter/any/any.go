// Package any provides between `vm.Value`s and built-in types.
package any

import (
	"fmt"

	"github.com/genkami/watson/pkg/vm"
)

// FromValue converts vm.Value into one of the following type:
// * int64
// * uint64
// * string
// * bool
// * (interface{})(nil)
// * map[string]interface{} (the value of which is also one or many of these types)
func FromValue(val *vm.Value) interface{} {
	switch val.Kind {
	case vm.KInt:
		return val.Int
	case vm.KUint:
		return val.Uint
	case vm.KFloat:
		return val.Float
	case vm.KString:
		return string(val.String)
	case vm.KObject:
		obj := map[string]interface{}{}
		for k, v := range val.Object {
			obj[k] = FromValue(v)
		}
		return obj
	case vm.KArray:
		arr := make([]interface{}, 0, len(val.Array))
		for _, v := range val.Array {
			arr = append(arr, FromValue(v))
		}
		return arr
	case vm.KBool:
		return val.Bool
	case vm.KNil:
		return nil
	default:
		panic(fmt.Errorf("invalid kind: %d", val.Kind))
	}
}

func ToValue(v interface{}) *vm.Value {
	if v == nil {
		return vm.NewNilValue()
	}
	switch v := v.(type) {
	case bool:
		return vm.NewBoolValue(v)
	case int:
		return vm.NewIntValue(int64(v))
	case int8:
		return vm.NewIntValue(int64(v))
	case int16:
		return vm.NewIntValue(int64(v))
	case int32:
		return vm.NewIntValue(int64(v))
	case int64:
		return vm.NewIntValue(v)
	case uint:
		return vm.NewUintValue(uint64(v))
	case uint8:
		return vm.NewUintValue(uint64(v))
	case uint16:
		return vm.NewUintValue(uint64(v))
	case uint32:
		return vm.NewUintValue(uint64(v))
	case uint64:
		return vm.NewUintValue(uint64(v))
	case []byte:
		w := make([]byte, len(v))
		copy(w, v)
		return vm.NewStringValue(w)
	case string:
		return vm.NewStringValue([]byte(v))
	case float32:
		return vm.NewFloatValue(float64(v))
	case float64:
		return vm.NewFloatValue(v)
	}
	panic(fmt.Errorf("can't convert %#v (%T) to *vm.Value", v, v))
}
