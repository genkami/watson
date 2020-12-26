package types

import (
	"fmt"
)

// ToGoObject converts Value into one of the following type:
//   * int64
//   * uint64
//   * float64
//   * string
//   * bool
//   * (interface{})(nil)
//   * []interface{} (the element of which is also one or many of these types)
//   * map[string]interface{} (the value of which is also one or many of these types)
func (val *Value) ToGoObject() interface{} {
	switch val.Kind {
	case Int:
		return val.Int
	case Uint:
		return val.Uint
	case Float:
		return val.Float
	case String:
		return string(val.String)
	case Object:
		obj := map[string]interface{}{}
		for k, v := range val.Object {
			obj[k] = v.ToGoObject()
		}
		return obj
	case Array:
		arr := make([]interface{}, 0, len(val.Array))
		for _, v := range val.Array {
			arr = append(arr, v.ToGoObject())
		}
		return arr
	case Bool:
		return val.Bool
	case Nil:
		return nil
	default:
		panic(fmt.Errorf("invalid kind: %d", val.Kind))
	}
}
