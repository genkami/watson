package watson_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson"
	"github.com/genkami/watson/pkg/types"
)

type User struct {
	FullName string `watson:"fullName"`
	Age      int    `watson:"age"`
}

type Department struct {
	Name    *DepartmentName
	Manager *User
}

type DepartmentName struct {
	Value string
}

func (d *DepartmentName) MarshalWatson() *types.Value {
	return types.NewStringValue([]byte(d.Value))
}

var _ types.Marshaler = &DepartmentName{}

func (d *DepartmentName) UnmarshalWatson(v *types.Value) error {
	if v.Kind != types.String {
		return fmt.Errorf("expected string but got %#v", v.Kind)
	}
	d.Value = string(v.String)
	return nil
}

func TestEncodeAndDecode(t *testing.T) {
	want := Department{
		Name: &DepartmentName{
			Value: "marketing",
		},
		Manager: &User{
			FullName: "Tanaka Taro",
			Age:      41,
		},
	}
	var got Department
	err := encodeThenDecode(&want, &got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(&want, &got); diff != "" {
		t.Fatalf("expected %#v but got %#v", &want, &got)
	}
}

func encodeThenDecode(in interface{}, out interface{}) error {
	var err error
	buf := bytes.NewBuffer(nil)
	enc := watson.NewEncoder(buf)
	err = enc.Encode(in)
	if err != nil {
		return err
	}
	dec := watson.NewDecoder(bytes.NewReader(buf.Bytes()))
	err = dec.Decode(out)
	if err != nil {
		return err
	}
	return nil
}
