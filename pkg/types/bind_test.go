package types_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/types"
)

func TestBindConvertsInt(t *testing.T) {
	var err error
	var got int
	var val = types.NewIntValue(123)
	var want int = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt8(t *testing.T) {
	var err error
	var got int8
	var val = types.NewIntValue(123)
	var want int8 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt16(t *testing.T) {
	var err error
	var got int16
	var val = types.NewIntValue(123)
	var want int16 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt32(t *testing.T) {
	var err error
	var got int32
	var val = types.NewIntValue(123)
	var want int32 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsInt64(t *testing.T) {
	var err error
	var got int64
	var val = types.NewIntValue(123)
	var want int64 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt(t *testing.T) {
	var err error
	var got int
	var val = types.NewIntValue(123)
	var want int = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt8(t *testing.T) {
	var err error
	var got int8
	var val = types.NewIntValue(123)
	var want int8 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt16(t *testing.T) {
	var err error
	var got int16
	var val = types.NewIntValue(123)
	var want int16 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt32(t *testing.T) {
	var err error
	var got int32
	var val = types.NewIntValue(123)
	var want int32 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsInt64(t *testing.T) {
	var err error
	var got int64
	var val = types.NewIntValue(123)
	var want int64 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
