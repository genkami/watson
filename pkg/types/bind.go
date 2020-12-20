package types

import (
	"fmt"
	"reflect"
)

func (v *Value) Bind(to interface{}) error {
	switch to := to.(type) {
	case *int:
		return bindInt(v, to)
	case *int8:
		return bindInt8(v, to)
	case *int16:
		return bindInt16(v, to)
	case *int32:
		return bindInt32(v, to)
	case *int64:
		return bindInt64(v, to)
	}
	return v.BindByReflection(reflect.ValueOf(to))
}

func bindInt(v *Value, to *int) error {
	if v.Kind != Int {
		return typeMismatch(v, Int)
	}
	*to = int(v.Int)
	return nil
}

func bindInt8(v *Value, to *int8) error {
	if v.Kind != Int {
		return typeMismatch(v, Int)
	}
	*to = int8(v.Int)
	return nil
}

func bindInt16(v *Value, to *int16) error {
	if v.Kind != Int {
		return typeMismatch(v, Int)
	}
	*to = int16(v.Int)
	return nil
}

func bindInt32(v *Value, to *int32) error {
	if v.Kind != Int {
		return typeMismatch(v, Int)
	}
	*to = int32(v.Int)
	return nil
}

func bindInt64(v *Value, to *int64) error {
	if v.Kind != Int {
		return typeMismatch(v, Int)
	}
	*to = int64(v.Int)
	return nil
}

func (v *Value) BindByReflection(to reflect.Value) error {
	if isPtr(to) {
		return bindPtrByReflection(v, to)
	}
	return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
}

func bindPtrByReflection(v *Value, to reflect.Value) error {
	switch to.Type().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		to.Elem().SetInt(v.Int)
	default:
		return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
	}
	return nil
}

func typeMismatch(v *Value, k Kind) error {
	return fmt.Errorf("cn't convert %#v to %#v", v.Kind, k)
}
