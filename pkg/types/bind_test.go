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

func TestBindConvertsNilInterface(t *testing.T) {
	var err error
	var got interface{}
	var val = types.NewNilValue()
	var want interface{} = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilPtr(t *testing.T) {
	var err error
	var got *int
	var val = types.NewNilValue()
	var want *int = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilSlice(t *testing.T) {
	var err error
	var got []int
	var val = types.NewNilValue()
	var want []int = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsNilMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewNilValue()
	var want map[string]int = nil
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
	})
	var want map[string]int = map[string]int{
		"hoge": 123,
	}
	err = val.Bind(&got)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindConvertsHeteroMap(t *testing.T) {
	var err error
	var got map[string]interface{}
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
		"fuga": types.NewStringValue([]byte("foo")),
		"bar": types.NewObjectValue(map[string]*types.Value{
			"baz": types.NewStringValue([]byte("quux")),
		}),
	})
	var want map[string]interface{} = map[string]interface{}{
		"hoge": int64(123),
		"fuga": "foo",
		"bar": map[string]interface{}{
			"baz": "quux",
		},
	}
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

func TestBindByReflectionConvertsNilInterface(t *testing.T) {
	var err error
	var got interface{}
	var val = types.NewNilValue()
	var want interface{} = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilPtr(t *testing.T) {
	var err error
	var got *int
	var val = types.NewNilValue()
	var want *int = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilSlice(t *testing.T) {
	var err error
	var got []int
	var val = types.NewNilValue()
	var want []int = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsNilMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewNilValue()
	var want map[string]int = nil
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsMap(t *testing.T) {
	var err error
	var got map[string]int
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
	})
	var want map[string]int = map[string]int{
		"hoge": 123,
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestBindByReflectionConvertsHeteroMap(t *testing.T) {
	var err error
	var got map[string]interface{}
	var val = types.NewObjectValue(map[string]*types.Value{
		"hoge": types.NewIntValue(123),
		"fuga": types.NewStringValue([]byte("foo")),
		"bar": types.NewObjectValue(map[string]*types.Value{
			"baz": types.NewStringValue([]byte("quux")),
		}),
	})
	var want map[string]interface{} = map[string]interface{}{
		"hoge": int64(123),
		"fuga": "foo",
		"bar": map[string]interface{}{
			"baz": "quux",
		},
	}
	err = val.BindByReflection(reflect.ValueOf(&got))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
