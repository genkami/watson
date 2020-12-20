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
	case *uint:
		return bindUint(v, to)
	case *uint8:
		return bindUint8(v, to)
	case *uint16:
		return bindUint16(v, to)
	case *uint32:
		return bindUint32(v, to)
	case *uint64:
		return bindUint64(v, to)
	case *float32:
		return bindFloat32(v, to)
	case *float64:
		return bindFloat64(v, to)
	case *string:
		return bindString(v, to)
	case *bool:
		return bindBool(v, to)
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

func bindUint(v *Value, to *uint) error {
	if v.Kind != Uint {
		return typeMismatch(v, Uint)
	}
	*to = uint(v.Uint)
	return nil
}

func bindUint8(v *Value, to *uint8) error {
	if v.Kind != Uint {
		return typeMismatch(v, Uint)
	}
	*to = uint8(v.Uint)
	return nil
}

func bindUint16(v *Value, to *uint16) error {
	if v.Kind != Uint {
		return typeMismatch(v, Uint)
	}
	*to = uint16(v.Uint)
	return nil
}

func bindUint32(v *Value, to *uint32) error {
	if v.Kind != Uint {
		return typeMismatch(v, Uint)
	}
	*to = uint32(v.Uint)
	return nil
}

func bindUint64(v *Value, to *uint64) error {
	if v.Kind != Uint {
		return typeMismatch(v, Uint)
	}
	*to = uint64(v.Uint)
	return nil
}

func bindFloat32(v *Value, to *float32) error {
	if v.Kind != Float {
		return typeMismatch(v, Float)
	}
	*to = float32(v.Float)
	return nil
}

func bindFloat64(v *Value, to *float64) error {
	if v.Kind != Float {
		return typeMismatch(v, Float)
	}
	*to = float64(v.Float)
	return nil
}

func bindString(v *Value, to *string) error {
	if v.Kind != String {
		return typeMismatch(v, String)
	}
	*to = string(v.String)
	return nil
}

func bindBool(v *Value, to *bool) error {
	if v.Kind != Bool {
		return typeMismatch(v, Bool)
	}
	*to = v.Bool
	return nil
}

func (v *Value) BindByReflection(to reflect.Value) error {
	if !isPtr(to) {
		return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
	}
	switch to.Type().Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.bindIntByReflection(to)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.bindUintByReflection(to)
	case reflect.Float32, reflect.Float64:
		return v.bindFloatByReflection(to)
	case reflect.String:
		return v.bindStringByReflection(to)
	case reflect.Bool:
		return v.bindBoolByReflection(to)
	case reflect.Ptr, reflect.Interface:
		return v.bindPtrByReflection(to)
	case reflect.Slice:
		return v.bindSliceByReflection(to)
	case reflect.Map:
		return v.bindMapByReflection(to)
	default:
		return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
	}
}

func (v *Value) bindIntByReflection(to reflect.Value) error {
	if v.Kind != Int {
		return typeMismatch(v, Int)
	}
	to.Elem().SetInt(v.Int)
	return nil
}

func (v *Value) bindUintByReflection(to reflect.Value) error {
	if v.Kind != Uint {
		return typeMismatch(v, Uint)
	}
	to.Elem().SetUint(v.Uint)
	return nil
}

func (v *Value) bindFloatByReflection(to reflect.Value) error {
	if v.Kind != Float {
		return typeMismatch(v, Float)
	}
	to.Elem().SetFloat(v.Float)
	return nil
}

func (v *Value) bindStringByReflection(to reflect.Value) error {
	if v.Kind != String {
		return typeMismatch(v, String)
	}
	to.Elem().SetString(string(v.String))
	return nil
}

func (v *Value) bindBoolByReflection(to reflect.Value) error {
	if v.Kind != Bool {
		return typeMismatch(v, Bool)
	}
	to.Elem().SetBool(bool(v.Bool))
	return nil
}

func (v *Value) bindSliceByReflection(to reflect.Value) error {
	if v.Kind == Nil {
		to.Elem().Set(reflect.Zero(to.Type().Elem()))
		return nil
	}
	return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
}

func (v *Value) bindMapByReflection(to reflect.Value) error {
	if v.Kind == Nil {
		to.Elem().Set(reflect.Zero(to.Type().Elem()))
		return nil
	}
	return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
}

func (v *Value) bindPtrByReflection(to reflect.Value) error {
	if v.Kind == Nil {
		to.Elem().Set(reflect.Zero(to.Type().Elem()))
		return nil
	}
	return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
}

func typeMismatch(v *Value, k Kind) error {
	return fmt.Errorf("cn't convert %#v to %#v", v.Kind, k)
}
