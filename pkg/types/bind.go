package types

import (
	"fmt"
	"reflect"
)

func (v *Value) Bind(to interface{}) error {
	return v.bind(to, newRootPath())
}

func (v *Value) bind(to interface{}, path path) error {
	switch to := to.(type) {
	case *int:
		return bindInt(v, to, path)
	case *int8:
		return bindInt8(v, to, path)
	case *int16:
		return bindInt16(v, to, path)
	case *int32:
		return bindInt32(v, to, path)
	case *int64:
		return bindInt64(v, to, path)
	case *uint:
		return bindUint(v, to, path)
	case *uint8:
		return bindUint8(v, to, path)
	case *uint16:
		return bindUint16(v, to, path)
	case *uint32:
		return bindUint32(v, to, path)
	case *uint64:
		return bindUint64(v, to, path)
	case *float32:
		return bindFloat32(v, to, path)
	case *float64:
		return bindFloat64(v, to, path)
	case *string:
		return bindString(v, to, path)
	case *bool:
		return bindBool(v, to, path)
	}
	if unmarshaler, ok := to.(Unmarshaler); ok {
		return unmarshaler.UnmarshalWatson(v)
	}
	return v.bindByReflection(reflect.ValueOf(to), path)
}

func bindInt(v *Value, to *int, path path) error {
	if v.Kind != Int {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = int(v.Int)
	return nil
}

func bindInt8(v *Value, to *int8, path path) error {
	if v.Kind != Int {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = int8(v.Int)
	return nil
}

func bindInt16(v *Value, to *int16, path path) error {
	if v.Kind != Int {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = int16(v.Int)
	return nil
}

func bindInt32(v *Value, to *int32, path path) error {
	if v.Kind != Int {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = int32(v.Int)
	return nil
}

func bindInt64(v *Value, to *int64, path path) error {
	if v.Kind != Int {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = int64(v.Int)
	return nil
}

func bindUint(v *Value, to *uint, path path) error {
	if v.Kind != Uint {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = uint(v.Uint)
	return nil
}

func bindUint8(v *Value, to *uint8, path path) error {
	if v.Kind != Uint {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = uint8(v.Uint)
	return nil
}

func bindUint16(v *Value, to *uint16, path path) error {
	if v.Kind != Uint {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = uint16(v.Uint)
	return nil
}

func bindUint32(v *Value, to *uint32, path path) error {
	if v.Kind != Uint {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = uint32(v.Uint)
	return nil
}

func bindUint64(v *Value, to *uint64, path path) error {
	if v.Kind != Uint {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = uint64(v.Uint)
	return nil
}

func bindFloat32(v *Value, to *float32, path path) error {
	if v.Kind != Float {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = float32(v.Float)
	return nil
}

func bindFloat64(v *Value, to *float64, path path) error {
	if v.Kind != Float {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = float64(v.Float)
	return nil
}

func bindString(v *Value, to *string, path path) error {
	if v.Kind != String {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = string(v.String)
	return nil
}

func bindBool(v *Value, to *bool, path path) error {
	if v.Kind != Bool {
		return &TypeMismatch{
			val:  v,
			t:    reflect.TypeOf(to).Elem(),
			path: path,
		}
	}
	*to = v.Bool
	return nil
}

func (v *Value) BindByReflection(to reflect.Value) error {
	return v.bindByReflection(to, newRootPath())
}

func (v *Value) bindByReflection(to reflect.Value, path path) error {
	if isUnmarshaler(to.Type()) {
		return v.bindToUnmarshalerByReflection(to, path)
	} else if isPtr(to) {
		return v.bindToPtrByReflection(to, path)
	}
	return fmt.Errorf("can't convert %#v to %s", v.Kind, to.Type().String())
}

func (v *Value) bindToUnmarshalerByReflection(to reflect.Value, path path) error {
	unmarshal := to.MethodByName("UnmarshalWatson")
	ret := unmarshal.Call([]reflect.Value{reflect.ValueOf(v)})[0].Interface()
	if err, ok := ret.(error); ok {
		return err
	}
	return nil
}

func (v *Value) bindToPtrByReflection(to reflect.Value, path path) error {
	casted, err := v.cast(to.Elem().Type(), path)
	if err != nil {
		return err
	}
	to.Elem().Set(casted)
	return nil
}

func (v *Value) cast(t reflect.Type, path path) (reflect.Value, error) {
	if isUnmarshaler(t) {
		return v.castToUnmarshaler(t, path)
	}
	switch t.Kind() {
	case reflect.Int:
		return v.castToInt(t, path)
	case reflect.Int8:
		return v.castToInt8(t, path)
	case reflect.Int16:
		return v.castToInt16(t, path)
	case reflect.Int32:
		return v.castToInt32(t, path)
	case reflect.Int64:
		return v.castToInt64(t, path)
	case reflect.Uint:
		return v.castToUint(t, path)
	case reflect.Uint8:
		return v.castToUint8(t, path)
	case reflect.Uint16:
		return v.castToUint16(t, path)
	case reflect.Uint32:
		return v.castToUint32(t, path)
	case reflect.Uint64:
		return v.castToUint64(t, path)
	case reflect.Float32:
		return v.castToFloat32(t, path)
	case reflect.Float64:
		return v.castToFloat64(t, path)
	case reflect.String:
		return v.castToString(t, path)
	case reflect.Bool:
		return v.castToBool(t, path)
	case reflect.Ptr:
		return v.castToPtr(t, path)
	case reflect.Interface:
		return v.castToInterface(t, path)
	case reflect.Slice:
		return v.castToSlice(t, path)
	case reflect.Array:
		return v.castToArray(t, path)
	case reflect.Map:
		return v.castToMap(t, path)
	case reflect.Struct:
		return v.castToStruct(t, path)
	default:
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
}

func (v *Value) castToInt(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(int(v.Int)), nil
}

func (v *Value) castToInt8(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(int8(v.Int)), nil
}

func (v *Value) castToInt16(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(int16(v.Int)), nil
}

func (v *Value) castToInt32(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(int32(v.Int)), nil
}

func (v *Value) castToInt64(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Int {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(int64(v.Int)), nil
}

func (v *Value) castToUint(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(uint(v.Uint)), nil
}

func (v *Value) castToUint8(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(uint8(v.Uint)), nil
}

func (v *Value) castToUint16(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(uint16(v.Uint)), nil
}

func (v *Value) castToUint32(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(uint32(v.Uint)), nil
}

func (v *Value) castToUint64(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Uint {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(uint64(v.Uint)), nil
}

func (v *Value) castToFloat32(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Float {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(float32(v.Float)), nil
}

func (v *Value) castToFloat64(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Float {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(v.Float), nil
}

func (v *Value) castToString(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != String {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(string(v.String)), nil
}

func (v *Value) castToBool(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Bool {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	return reflect.ValueOf(v.Bool), nil
}

func (v *Value) castToSlice(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind == Nil {
		return reflect.Zero(t), nil
	}
	if v.Kind == Array {
		arr := reflect.MakeSlice(t, len(v.Array), len(v.Array))
		err := v.setToArray(arr, path)
		if err != nil {
			return reflect.Value{}, err
		}
		return arr, nil
	}
	return reflect.Value{}, &TypeMismatch{
		val:  v,
		t:    t,
		path: path,
	}
}

func (v *Value) castToArray(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Array {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	if len(v.Array) > t.Len() {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	parr := reflect.New(t)
	err := v.setToArray(parr.Elem(), path)
	if err != nil {
		return reflect.Value{}, err
	}
	return parr.Elem(), nil
}

func (v *Value) setToArray(arr reflect.Value, path path) error {
	t := arr.Type()
	if v.Kind != Array || !(t.Kind() == reflect.Slice || t.Kind() == reflect.Array) {
		return &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	elemType := t.Elem()
	for i, e := range v.Array {
		elem, err := e.cast(elemType, newIndexPath(path, i))
		if err != nil {
			return err
		}
		arr.Index(i).Set(elem)
	}
	return nil
}

func (v *Value) castToMap(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind == Nil {
		return reflect.Zero(t), nil
	}
	if v.Kind == Object {
		obj := reflect.MakeMap(t)
		err := v.addToMap(obj, path)
		if err != nil {
			return reflect.Value{}, err
		}
		return obj, nil
	}
	return reflect.Value{}, &TypeMismatch{
		val:  v,
		t:    t,
		path: path,
	}
}

func (v *Value) addToMap(obj reflect.Value, path path) error {
	t := obj.Type()
	if v.Kind != Object || t.Kind() != reflect.Map {
		return &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	keyType := t.Key()
	elemType := t.Elem()
	if keyType.Kind() != reflect.String {
		return &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	for k, e := range v.Object {
		key := reflect.ValueOf(k)
		elem, err := e.cast(elemType, newFieldPath(path, k))
		if err != nil {
			return err
		}
		obj.SetMapIndex(key, elem)
	}
	return nil
}

func (v *Value) castToPtr(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind == Nil {
		return reflect.Zero(t), nil
	}
	elem, err := v.cast(t.Elem(), path)
	if err != nil {
		return reflect.Value{}, err
	}
	ptr := reflect.New(t.Elem())
	ptr.Elem().Set(elem)
	return ptr, nil
}

func (v *Value) castToInterface(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind == Nil {
		return reflect.Zero(t), nil
	}
	var any interface{}
	if t == reflect.TypeOf(&any).Elem() {
		return reflect.ValueOf(v.ToGoObject()), nil
	}
	return reflect.Value{}, &TypeMismatch{
		val:  v,
		t:    t,
		path: path,
	}
}

func (v *Value) castToStruct(t reflect.Type, path path) (reflect.Value, error) {
	if v.Kind != Object {
		return reflect.Value{}, &TypeMismatch{
			val:  v,
			t:    t,
			path: path,
		}
	}
	pobj := reflect.New(t)
	obj := pobj.Elem()
	for k, v := range v.Object {
		tag, ok := findField(k, obj)
		if !ok {
			continue
		}
		if tag.ShouldAlwaysOmit() {
			continue
		}
		field := tag.FieldOf(obj)
		err := v.BindByReflection(field.Addr())
		if err != nil {
			return reflect.Value{}, err
		}
	}
	for _, tag := range inlineFields(obj) {
		field := tag.FieldOf(obj)
		err := v.BindByReflection(field.Addr())
		if err != nil {
			return reflect.Value{}, err
		}
	}
	return obj, nil

}

func (v *Value) castToUnmarshaler(t reflect.Type, path path) (reflect.Value, error) {
	var obj reflect.Value
	if t.Kind() == reflect.Ptr {
		obj = reflect.New(t.Elem())
	} else {
		obj = reflect.New(t).Elem()
	}
	err := v.bindToUnmarshalerByReflection(obj, path)
	if err != nil {
		return reflect.Value{}, err
	}
	return obj, nil
}
