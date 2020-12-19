package types

import (
	"reflect"
	"strings"
)

const (
	tagId = "watson"
)

type tag struct {
	name string
	f    *reflect.StructField
}

func parseTag(f *reflect.StructField) *tag {
	name := f.Tag.Get(tagId)
	return &tag{name: name, f: f}
}

func (t *tag) key() string {
	if t.name == "" {
		return strings.ToLower(t.f.Name)
	}
	return t.name
}
