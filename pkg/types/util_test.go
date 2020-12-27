package types_test

import (
	"fmt"

	"github.com/genkami/watson/pkg/types"
)

type untagged struct {
	Name     string
	LongName string
}

type nested struct {
	Value *nestedInner
}

type nestedInner struct {
	Value int
}

type embedded struct {
	Field int
	EmbeddedInner
}

type EmbeddedInner struct {
	AnotherField int
}

type tagged struct {
	Field int `watson:"customName"`
}

type private struct {
	PublicField  int
	privateField int
}

type omitempty struct {
	Field1 *int `watson:",omitempty"`
	Field2 *int `watson:"field2,omitempty"`
}

type alwaysomit struct {
	ShouldBeIncluded int `watson:"shouldBeIncluded"`
	ShouldBeOmitted  int `watson:"-"`
}

type inline struct {
	Field int
	Inner inlineInner `watson:",inline"`
}

type inlineInner struct {
	NestedField int
}

type customMarshaler struct {
	SomeField int
}

func (m *customMarshaler) MarshalWatson() *types.Value {
	return types.NewObjectValue(map[string]*types.Value{
		"custom":          types.NewBoolValue(true),
		"customFieldName": types.NewIntValue(int64(m.SomeField)),
	})
}

type primitiveMarshaler int

func (p primitiveMarshaler) MarshalWatson() *types.Value {
	return types.NewStringValue([]byte(fmt.Sprintf("value=%d", p)))
}

type customUnmarshalerOuter struct {
	Unmarshaler *customUnmarshaler
}

type customUnmarshaler struct {
	SomeField int
}

func (u *customUnmarshaler) UnmarshalWatson(v *types.Value) error {
	if v.Kind != types.Object {
		return fmt.Errorf("value is not an Object")
	}
	k, ok := v.Object["customKey"]
	if !ok {
		return fmt.Errorf("value does not have customKey")
	}
	if k.Kind != types.Int {
		return fmt.Errorf("custom key is not an int")
	}
	u.SomeField = int(k.Int)
	return nil
}

type primitiveUnmarshaler int

func (p *primitiveUnmarshaler) UnmarshalWatson(v *types.Value) error {
	if v.Kind != types.Object {
		return fmt.Errorf("value is not an Object")
	}
	k, ok := v.Object["customKey"]
	if !ok {
		return fmt.Errorf("value does not have customKey")
	}
	if k.Kind != types.Int {
		return fmt.Errorf("custom key is not an int")
	}
	*p = primitiveUnmarshaler(k.Int)
	return nil
}
