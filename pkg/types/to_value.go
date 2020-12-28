package types

import (
	"fmt"
	"reflect"
)

// ToValue converts an arbitrary value into *Value.
//
// See watson.Marshal for details.
func ToValue(v interface{}) (*Value, error) {
	if v == nil {
		return NewNilValue(), nil
	}
	switch v := v.(type) {
	case bool:
		return NewBoolValue(v), nil
	case int:
		return NewIntValue(int64(v)), nil
	case int8:
		return NewIntValue(int64(v)), nil
	case int16:
		return NewIntValue(int64(v)), nil
	case int32:
		return NewIntValue(int64(v)), nil
	case int64:
		return NewIntValue(v), nil
	case uint:
		return NewUintValue(uint64(v)), nil
	case uint8:
		return NewUintValue(uint64(v)), nil
	case uint16:
		return NewUintValue(uint64(v)), nil
	case uint32:
		return NewUintValue(uint64(v)), nil
	case uint64:
		return NewUintValue(uint64(v)), nil
	case string:
		return NewStringValue([]byte(v)), nil
	case float32:
		return NewFloatValue(float64(v)), nil
	case float64:
		return NewFloatValue(v), nil
	}
	if marshaler, ok := v.(Marshaler); ok {
		return marshaler.MarshalWatson()
	}
	vv := reflect.ValueOf(v)
	return ToValueByReflection(vv)
}

// `ToValueByReflection` does almost the same thing as `ToValue`, but it always uses reflection.
func ToValueByReflection(v reflect.Value) (*Value, error) {
	if isMarshaler(v) {
		return marshalerToValueByReflection(v)
	} else if isIntFamily(v) {
		return intToValueByReflection(v)
	} else if isUintFamily(v) {
		return uintToValueByReflection(v)
	} else if isFloatFamily(v) {
		return floatToValueByReflection(v)
	} else if isBool(v) {
		return boolToValueByReflection(v)
	} else if isString(v) {
		return stringToValueByReflection(v)
	} else if isArray(v) {
		return sliceOrArrayToValueByReflection(v)
	} else if isStruct(v) {
		return structToValueByReflection(v)
	} else if isNil(v) {
		// Marshalers should be placed before nil so as to handle `MarshalWatson` correctly.
		return NewNilValue(), nil
		// Maps, slices, and pointers should be placed after nil so as to convert nil into Nil correctly.
	} else if isPtr(v) {
		return ptrToValueByReflection(v)
	} else if isMap(v) {
		return mapToValueByReflection(v)
	} else if isSlice(v) {
		return sliceOrArrayToValueByReflection(v)
	}

	return nil, fmt.Errorf("can't convert %s to *Value", v.Type().String())
}

func intToValueByReflection(v reflect.Value) (*Value, error) {
	return NewIntValue(v.Int()), nil
}

func uintToValueByReflection(v reflect.Value) (*Value, error) {
	return NewUintValue(v.Uint()), nil
}

func floatToValueByReflection(v reflect.Value) (*Value, error) {
	return NewFloatValue(v.Float()), nil
}

func boolToValueByReflection(v reflect.Value) (*Value, error) {
	return NewBoolValue(v.Bool()), nil
}

func stringToValueByReflection(v reflect.Value) (*Value, error) {
	return NewStringValue([]byte(v.String())), nil
}

func mapToValueByReflection(v reflect.Value) (*Value, error) {
	var err error
	obj := map[string]*Value{}
	iter := v.MapRange()
	for iter.Next() {
		key := iter.Key()
		k, ok := key.Interface().(string)
		if !ok {
			return nil, fmt.Errorf("can't convert %s to string", key.Type().String())
		}
		elem := iter.Value()
		var elemVal *Value
		if elem.CanInterface() {
			elemVal, err = ToValue(elem.Interface())
		} else {
			elemVal, err = ToValueByReflection(elem)
		}
		if err != nil {
			return nil, err
		}
		obj[k] = elemVal
	}
	return NewObjectValue(obj), nil
}

func sliceOrArrayToValueByReflection(v reflect.Value) (*Value, error) {
	var err error
	arr := []*Value{}
	size := v.Len()
	for i := 0; i < size; i++ {
		elem := v.Index(i)
		var elemVal *Value
		if elem.CanInterface() {
			elemVal, err = ToValue(elem.Interface())
		} else {
			elemVal, err = ToValueByReflection(elem)
		}
		if err != nil {
			return nil, err
		}
		arr = append(arr, elemVal)
	}
	return NewArrayValue(arr), nil
}

func ptrToValueByReflection(v reflect.Value) (*Value, error) {
	elem := v.Elem()
	if elem.CanInterface() {
		return ToValue(elem.Interface())
	} else {
		return ToValueByReflection(elem)
	}
}

func structToValueByReflection(v reflect.Value) (*Value, error) {
	obj := map[string]*Value{}
	err := addFields(obj, v)
	if err != nil {
		return nil, err
	}
	return NewObjectValue(obj), nil
}

func addFields(obj map[string]*Value, v reflect.Value) error {
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
			err := addFields(obj, elem)
			if err != nil {
				return err
			}
		} else if elem.CanInterface() {
			elemVal, err := ToValue(elem.Interface())
			if err != nil {
				return err
			}
			obj[name] = elemVal
		} else {
			elemVal, err := ToValueByReflection(elem)
			if err != nil {
				return err
			}
			obj[name] = elemVal
		}
	}
	return nil
}

func marshalerToValueByReflection(v reflect.Value) (*Value, error) {
	marshal := v.MethodByName("MarshalWatson")
	ret := marshal.Call([]reflect.Value{})
	val := ret[0].Interface().(*Value)
	if err, ok := ret[1].Interface().(error); ok {
		return val, err
	}
	return val, nil
}
