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
	return v.bindByReflection(to.Elem())
}

func (v *Value) bindByReflection(to reflect.Value) error {
	casted, err := v.castByReflection(to.Type())
	if err != nil {
		return err
	}
	to.Set(casted)
	return nil
}

func (v *Value) castByReflection(t reflect.Type) (reflect.Value, error) {
	switch t.Kind() {
	case reflect.Int:
		return v.castToInt(t)
	case reflect.Int8:
		return v.castToInt8(t)
	case reflect.Int16:
		return v.castToInt16(t)
	case reflect.Int32:
		return v.castToInt32(t)
	case reflect.Int64:
		return v.castToInt64(t)
	case reflect.Uint:
		return v.castToUint(t)
	case reflect.Uint8:
		return v.castToUint8(t)
	case reflect.Uint16:
		return v.castToUint16(t)
	case reflect.Uint32:
		return v.castToUint32(t)
	case reflect.Uint64:
		return v.castToUint64(t)
	case reflect.Float32:
		return v.castToFloat32(t)
	case reflect.Float64:
		return v.castToFloat64(t)
	case reflect.String:
		return v.castToString(t)
	case reflect.Bool:
		return v.castToBool(t)
	case reflect.Ptr:
		return v.castToPtr(t)
	case reflect.Interface:
		return v.castToInterface(t)
	case reflect.Slice:
		return v.castToSlice(t)
	case reflect.Map:
		return v.castToMap(t)
	default:
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
}

func (v *Value) castToInt(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(int(v.Int)), nil
}

func (v *Value) castToInt8(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(int8(v.Int)), nil
}

func (v *Value) castToInt16(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(int16(v.Int)), nil
}

func (v *Value) castToInt32(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(int32(v.Int)), nil
}

func (v *Value) castToInt64(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(int64(v.Int)), nil
}

func (v *Value) castToUint(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(uint(v.Uint)), nil
}

func (v *Value) castToUint8(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(uint8(v.Uint)), nil
}

func (v *Value) castToUint16(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(uint16(v.Uint)), nil
}

func (v *Value) castToUint32(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(uint32(v.Uint)), nil
}

func (v *Value) castToUint64(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(uint64(v.Uint)), nil
}

func (v *Value) castToFloat32(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Float {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(float32(v.Float)), nil
}

func (v *Value) castToFloat64(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Float {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(v.Float), nil
}

func (v *Value) castToString(t reflect.Type) (reflect.Value, error) {
	if v.Kind != String {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(string(v.String)), nil
}

func (v *Value) castToBool(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Bool {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.ValueOf(v.Bool), nil
}

func (v *Value) castToSlice(t reflect.Type) (reflect.Value, error) {
	if v.Kind == Nil {
		return reflect.Zero(t), nil
	}
	return reflect.Value{}, typeMismatchByReflection(v, t)
}

func (v *Value) castToMap(t reflect.Type) (reflect.Value, error) {
	if v.Kind == Nil {
		return reflect.Zero(t), nil
	}
	return reflect.Value{}, typeMismatchByReflection(v, t)
}

func (v *Value) castToPtr(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Nil {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.Zero(t), nil
}

func (v *Value) castToInterface(t reflect.Type) (reflect.Value, error) {
	if v.Kind != Nil {
		return reflect.Value{}, typeMismatchByReflection(v, t)
	}
	return reflect.Zero(t), nil
}

func typeMismatch(v *Value, k Kind) error {
	return fmt.Errorf("cn't convert %#v to %#v", v.Kind, k)
}

func typeMismatchByReflection(v *Value, t reflect.Type) error {
	return fmt.Errorf("can't convert %#v to %s", v.Kind, t.String())
}
