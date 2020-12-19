// Package any provides between `vm.Value`s and built-in types.
package any

import (
	"fmt"
	"reflect"

	"github.com/genkami/watson/pkg/types"
)

// FromValue converts types.Value into one of the following type:
// * int64
// * uint64
// * string
// * bool
// * (interface{})(nil)
// * map[string]interface{} (the value of which is also one or many of these types)
func FromValue(val *types.Value) interface{} {
	switch val.Kind {
	case types.Int:
		return val.Int
	case types.Uint:
		return val.Uint
	case types.Float:
		return val.Float
	case types.String:
		return string(val.String)
	case types.Object:
		obj := map[string]interface{}{}
		for k, v := range val.Object {
			obj[k] = FromValue(v)
		}
		return obj
	case types.Array:
		arr := make([]interface{}, 0, len(val.Array))
		for _, v := range val.Array {
			arr = append(arr, FromValue(v))
		}
		return arr
	case types.Bool:
		return val.Bool
	case types.Nil:
		return nil
	default:
		panic(fmt.Errorf("invalid kind: %d", val.Kind))
	}
}

func ToValue(v interface{}) *types.Value {
	if v == nil {
		return types.NewNilValue()
	}
	switch v := v.(type) {
	case bool:
		return types.NewBoolValue(v)
	case int:
		return types.NewIntValue(int64(v))
	case int8:
		return types.NewIntValue(int64(v))
	case int16:
		return types.NewIntValue(int64(v))
	case int32:
		return types.NewIntValue(int64(v))
	case int64:
		return types.NewIntValue(v)
	case uint:
		return types.NewUintValue(uint64(v))
	case uint8:
		return types.NewUintValue(uint64(v))
	case uint16:
		return types.NewUintValue(uint64(v))
	case uint32:
		return types.NewUintValue(uint64(v))
	case uint64:
		return types.NewUintValue(uint64(v))
	case string:
		return types.NewStringValue([]byte(v))
	case float32:
		return types.NewFloatValue(float64(v))
	case float64:
		return types.NewFloatValue(v)
	}
	vv := reflect.ValueOf(v)
	return reflectValueToValue(vv)
}

func reflectValueToValue(v reflect.Value) *types.Value {
	if isIntFamily(v) {
		return reflectIntToValue(v)
	} else if isUintFamily(v) {
		return reflectUintToValue(v)
	} else if isFloatFamily(v) {
		return reflectFloatToValue(v)
	} else if isBool(v) {
		return reflectBoolToValue(v)
	} else if isString(v) {
		return reflectStringToValue(v)
	} else if v.IsNil() {
		// Marshalers should be placed before nil so as to handle `MarshalWatson` correctly.
		return types.NewNilValue()
		// Maps, slices, and structs should be placed after nil so as to convert nil into Nil correctly.
	} else if isMapConvertibleToValue(v) {
		return reflectMapToValue(v)
	} else if isSlice(v) {
		return reflectSliceToValue(v)
	}

	panic(fmt.Errorf("can't convert %s to *types.Value", v.Type().String()))
}

func reflectIntToValue(v reflect.Value) *types.Value {
	return types.NewIntValue(v.Int())
}

func reflectUintToValue(v reflect.Value) *types.Value {
	return types.NewUintValue(v.Uint())
}

func reflectFloatToValue(v reflect.Value) *types.Value {
	return types.NewFloatValue(v.Float())
}

func reflectBoolToValue(v reflect.Value) *types.Value {
	return types.NewBoolValue(v.Bool())
}

func reflectStringToValue(v reflect.Value) *types.Value {
	return types.NewStringValue([]byte(v.String()))
}

func reflectMapToValue(v reflect.Value) *types.Value {
	obj := map[string]*types.Value{}
	iter := v.MapRange()
	for iter.Next() {
		k := iter.Key().String()
		v := iter.Value()
		if v.CanInterface() {
			obj[k] = ToValue(v.Interface())
		} else {
			obj[k] = reflectValueToValue(v)
		}
	}
	return types.NewObjectValue(obj)
}

func reflectSliceToValue(v reflect.Value) *types.Value {
	arr := []*types.Value{}
	size := v.Len()
	for i := 0; i < size; i++ {
		elem := v.Index(i)
		if elem.CanInterface() {
			arr = append(arr, ToValue(elem.Interface()))
		} else {
			arr = append(arr, reflectValueToValue(elem))
		}
	}
	return types.NewArrayValue(arr)
}

func isIntFamily(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}

func isUintFamily(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func isFloatFamily(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isBool(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Bool
}

func isString(v reflect.Value) bool {
	return v.Type().Kind() == reflect.String
}

func isMapConvertibleToValue(v reflect.Value) bool {
	t := v.Type()
	return t.Kind() == reflect.Map && t.Key().Kind() == reflect.String
}

func isSlice(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Slice
}
