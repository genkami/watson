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

func TestBindConvertsUint(t *testing.T) {
	var err error
	var got uint
	var val = types.NewUintValue(123)
	var want uint = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint8(t *testing.T) {
	var err error
	var got uint8
	var val = types.NewUintValue(123)
	var want uint8 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint16(t *testing.T) {
	var err error
	var got uint16
	var val = types.NewUintValue(123)
	var want uint16 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint32(t *testing.T) {
	var err error
	var got uint32
	var val = types.NewUintValue(123)
	var want uint32 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsUint64(t *testing.T) {
	var err error
	var got uint64
	var val = types.NewUintValue(123)
	var want uint64 = 123
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsFloat32(t *testing.T) {
	var err error
	var got float32
	var val = types.NewFloatValue(1.23e4)
	var want float32 = 1.23e4
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	if !closeEnough(float64(want), float64(got)) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestBindConvertsFloat64(t *testing.T) {
	var err error
	var got float64
	var val = types.NewFloatValue(1.23e45)
	var want float64 = 1.23e45
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsString(t *testing.T) {
	var err error
	var got string
	var val = types.NewStringValue([]byte("hoge"))
	var want string = "hoge"
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsBool(t *testing.T) {
	var err error
	var got bool
	var val = types.NewBoolValue(true)
	var want bool = true
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

func TestBindByReflectionConvertsUint(t *testing.T) {
	var err error
	var got uint
	var val = types.NewUintValue(123)
	var want uint = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint8(t *testing.T) {
	var err error
	var got uint8
	var val = types.NewUintValue(123)
	var want uint8 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint16(t *testing.T) {
	var err error
	var got uint16
	var val = types.NewUintValue(123)
	var want uint16 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint32(t *testing.T) {
	var err error
	var got uint32
	var val = types.NewUintValue(123)
	var want uint32 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsUint64(t *testing.T) {
	var err error
	var got uint64
	var val = types.NewUintValue(123)
	var want uint64 = 123
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsFloat32(t *testing.T) {
	var err error
	var got float32
	var val = types.NewFloatValue(1.23e4)
	var want float32 = 1.23e4
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
	if !closeEnough(float64(want), float64(got)) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestBindByReflectionConvertsFloat64(t *testing.T) {
	var err error
	var got float64
	var val = types.NewFloatValue(1.23e45)
	var want float64 = 1.23e45
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsString(t *testing.T) {
	var err error
	var got string
	var val = types.NewStringValue([]byte("hoge"))
	var want string = "hoge"
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsBool(t *testing.T) {
	var err error
	var got bool
	var val = types.NewBoolValue(true)
	var want bool = true
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
