package types

import (
	"fmt"
	"reflect"
)

func ToValue(v interface{}) *Value {
	if v == nil {
		return NewNilValue()
	}
	switch v := v.(type) {
	case bool:
		return NewBoolValue(v)
	case int:
		return NewIntValue(int64(v))
	case int8:
		return NewIntValue(int64(v))
	case int16:
		return NewIntValue(int64(v))
	case int32:
		return NewIntValue(int64(v))
	case int64:
		return NewIntValue(v)
	case uint:
		return NewUintValue(uint64(v))
	case uint8:
		return NewUintValue(uint64(v))
	case uint16:
		return NewUintValue(uint64(v))
	case uint32:
		return NewUintValue(uint64(v))
	case uint64:
		return NewUintValue(uint64(v))
	case string:
		return NewStringValue([]byte(v))
	case float32:
		return NewFloatValue(float64(v))
	case float64:
		return NewFloatValue(v)
	}
	vv := reflect.ValueOf(v)
	return ToValueByReflection(vv)
}

func ToValueByReflection(v reflect.Value) *Value {
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
	} else if isArray(v) {
		return reflectSliceOrArrayToValue(v)
	} else if isNil(v) {
		// Marshalers should be placed before nil so as to handle `MarshalWatson` correctly.
		return NewNilValue()
		// Maps, slices, and pointers should be placed after nil so as to convert nil into Nil correctly.
	} else if isPtr(v) {
		return reflectPtrToValue(v)
	} else if isMapConvertibleToValue(v) {
		return reflectMapToValue(v)
	} else if isSlice(v) {
		return reflectSliceOrArrayToValue(v)
	}

	panic(fmt.Errorf("can't convert %s to *Value", v.Type().String()))
}

func reflectIntToValue(v reflect.Value) *Value {
	return NewIntValue(v.Int())
}

func reflectUintToValue(v reflect.Value) *Value {
	return NewUintValue(v.Uint())
}

func reflectFloatToValue(v reflect.Value) *Value {
	return NewFloatValue(v.Float())
}

func reflectBoolToValue(v reflect.Value) *Value {
	return NewBoolValue(v.Bool())
}

func reflectStringToValue(v reflect.Value) *Value {
	return NewStringValue([]byte(v.String()))
}

func reflectMapToValue(v reflect.Value) *Value {
	obj := map[string]*Value{}
	iter := v.MapRange()
	for iter.Next() {
		k := iter.Key().String()
		v := iter.Value()
		if v.CanInterface() {
			obj[k] = ToValue(v.Interface())
		} else {
			obj[k] = ToValueByReflection(v)
		}
	}
	return NewObjectValue(obj)
}

func reflectSliceOrArrayToValue(v reflect.Value) *Value {
	arr := []*Value{}
	size := v.Len()
	for i := 0; i < size; i++ {
		elem := v.Index(i)
		if elem.CanInterface() {
			arr = append(arr, ToValue(elem.Interface()))
		} else {
			arr = append(arr, ToValueByReflection(elem))
		}
	}
	return NewArrayValue(arr)
}

func reflectPtrToValue(v reflect.Value) *Value {
	elem := v.Elem()
	if elem.CanInterface() {
		return ToValue(elem.Interface())
	} else {
		return ToValueByReflection(elem)
	}
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

func isNil(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func isMapConvertibleToValue(v reflect.Value) bool {
	t := v.Type()
	return t.Kind() == reflect.Map && t.Key().Kind() == reflect.String
}

func isSlice(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Slice
}

func isArray(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Array
}

func isPtr(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Ptr
}
