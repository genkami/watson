package types

import (
	"fmt"
	"reflect"
)

// ToValue converts an arbitrary value `v` into `Value` by applying the following rules:
//   * If v is any of int, int8, int16, int32, or int64, then v is converted to Int.
//   * If v is any of uint, uint8, uint16, uint32, or uint64, then v is converted to Uint.
//   * If v is float32 or float64, then v is converted to Float.
//   * If v is bool, then v is converted to Bool.
//   * If v is string, then v is converted to String.
//   * If v is a struct that implements `Marshaler`, then v is converted to Value by calling `v.MarshalWatson()`.
//   * If v is a struct that does not implement `Marshaler`, then v is converted to Object with its keys correspond to the fields of v.
//   * If v is a slice or an array, then v is converted to Array with its elements converted by these rules.
//   * If v is a map, then v is converted to Object with its elements converted by these rules.
//   * If v is a pointer, then v is converted to `Value` by converting `*v` with these rules.
//
// Note that you can configure struct fields by adding "watson" tag to fields.
// Tag must be like `watson:"name,flag1,flag2,...,flagN"`.
// If `ToValue` finds a field that has such tag, it uses `name` as a key of output instead of using the name of the field, or omits such field if `name` equals to "-".
//
// Currntly these flags are available:
//   omitempty      If the field is zero value, it will be omitted from the output.
//   inline         Inline the field. Currently the field must be a struct.
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
	if marshaler, ok := v.(Marshaler); ok {
		return marshaler.MarshalWatson()
	}
	vv := reflect.ValueOf(v)
	return ToValueByReflection(vv)
}

// `ToValueByReflection` does almost the same thing as `ToValue`, but it always uses reflection.
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
	} else if isMarshaler(v) {
		return reflectMarshalerToValue(v)
	} else if isStruct(v) {
		return reflectStructToValue(v)
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

func reflectStructToValue(v reflect.Value) *Value {
	obj := map[string]*Value{}
	addFields(obj, v)
	return NewObjectValue(obj)
}

func addFields(obj map[string]*Value, v reflect.Value) {
	size := v.NumField()
	t := v.Type()
	for i := 0; i < size; i++ {
		field := t.Field(i)
		tag := parseTag(&field)
		if tag.ShouldAlwaysOmit() {
			continue
		}
		name := tag.Key()
		elem := v.Field(i)
		if tag.OmitEmpty() && elem.IsZero() {
			continue
		}
		if tag.Inline() {
			addFields(obj, elem)
		} else if elem.CanInterface() {
			obj[name] = ToValue(elem.Interface())
		} else {
			obj[name] = ToValueByReflection(elem)
		}
	}
}

func reflectMarshalerToValue(v reflect.Value) *Value {
	marshal := v.MethodByName("MarshalWatson")
	ret := marshal.Call([]reflect.Value{})
	return ret[0].Interface().(*Value)
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

func isStruct(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Struct
}

func isMarshaler(v reflect.Value) bool {
	var marshaler Marshaler
	return v.Type().Implements(reflect.TypeOf(&marshaler).Elem())
}
