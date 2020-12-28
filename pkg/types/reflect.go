package types

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	tagId          = "watson"
	attrAlwaysOmit = "-"
	attrOmitEmpty  = "omitempty"
	attrInline     = "inline"
)

type tag struct {
	name       string
	f          *reflect.StructField
	omitempty  bool
	alwaysomit bool
	inline     bool
}

func findField(key string, obj reflect.Value) (*tag, bool) {
	t := obj.Type()
	size := obj.NumField()
	for i := 0; i < size; i++ {
		f := t.Field(i)
		tag := parseTag(&f)
		if tag.Key() == key {
			return tag, true
		}
	}
	return nil, false
}

func inlineFields(obj reflect.Value) []*tag {
	t := obj.Type()
	tags := make([]*tag, 0)
	size := obj.NumField()
	for i := 0; i < size; i++ {
		f := t.Field(i)
		tag := parseTag(&f)
		if tag.Inline() {
			tags = append(tags, tag)
		}
	}
	return tags
}

func parseTag(f *reflect.StructField) *tag {
	tag := &tag{f: f}
	name := f.Tag.Get(tagId)
	if name == "" {
		return tag
	}
	attrs := strings.Split(name, ",")
	first := true
	for _, attr := range attrs {
		if first {
			if attr == attrAlwaysOmit {
				tag.alwaysomit = true
			} else {
				tag.name = attr
			}
			first = false
			continue
		}
		switch attr {
		case attrOmitEmpty:
			tag.omitempty = true
		case attrInline:
			tag.inline = true
		}
	}
	return tag
}

func (t *tag) Key() string {
	if t.name == "" {
		return strings.ToLower(t.f.Name)
	}
	return t.name
}

func (t *tag) ShouldAlwaysOmit() bool {
	if t.alwaysomit {
		return true
	}
	r, _ := utf8.DecodeRuneInString(t.f.Name)
	return unicode.IsLower(r)
}

func (t *tag) OmitEmpty() bool {
	return t.omitempty
}

func (t *tag) Inline() bool {
	return t.inline
}

func (t *tag) FieldOf(v reflect.Value) reflect.Value {
	return v.FieldByIndex(t.f.Index)
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

func isMap(v reflect.Value) bool {
	t := v.Type()
	return t.Kind() == reflect.Map
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

func isUnmarshaler(t reflect.Type) bool {
	var unmarshaler Unmarshaler
	return t.Implements(reflect.TypeOf(&unmarshaler).Elem())
}

// TypeMismatch is an error that indicates that a given Value can't be converted into expected type.
type TypeMismatch struct {
	val  *Value
	t    reflect.Type
	path path
}

func (e *TypeMismatch) Error() string {
	return fmt.Sprintf("can't convert %#v to %s (at %s)",
		e.val.Kind, e.t.String(), e.path.string())
}
