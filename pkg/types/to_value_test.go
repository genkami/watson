package types_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/types"
)

func TestToValueConvertsNilInterface(t *testing.T) {
	want := types.NewNilValue()
	got, err := types.ToValue(nil)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilPointer(t *testing.T) {
	var p *int = nil
	want := types.NewNilValue()
	got, err := types.ToValue(p)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilSlice(t *testing.T) {
	var p []int = nil
	want := types.NewNilValue()
	got, err := types.ToValue(p)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilMap(t *testing.T) {
	var p map[string]interface{} = nil
	want := types.NewNilValue()
	got, err := types.ToValue(p)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsTrue(t *testing.T) {
	want := types.NewBoolValue(true)
	got, err := types.ToValue(true)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFalse(t *testing.T) {
	want := types.NewBoolValue(false)
	got, err := types.ToValue(false)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt(t *testing.T) {
	want := types.NewIntValue(-12345678)
	got, err := types.ToValue(int(-12345678))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt8(t *testing.T) {
	want := types.NewIntValue(64)
	got, err := types.ToValue(int8(64))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt16(t *testing.T) {
	want := types.NewIntValue(256)
	got, err := types.ToValue(int16(256))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt32(t *testing.T) {
	want := types.NewIntValue(65536)
	got, err := types.ToValue(int32(65536))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt64(t *testing.T) {
	want := types.NewIntValue(1234567)
	got, err := types.ToValue(int64(1234567))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValue(uint(12345))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint8(t *testing.T) {
	want := types.NewUintValue(255)
	got, err := types.ToValue(uint8(255))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint16(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValue(uint16(12345))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint32(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValue(uint32(12345))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint64(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValue(uint64(12345))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsString(t *testing.T) {
	want := types.NewStringValue([]byte("hogefuga"))
	got, err := types.ToValue("hogefuga")
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFloat32(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got, err := types.ToValue(float32(1.2345e-6))
	if err != nil {
		t.Fatal(err)
	}
	if got.Kind != types.Float {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueConvertsFloat64(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got, err := types.ToValue(float64(1.2345e-6))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsMap(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	got, err := types.ToValue(map[string]string{"hello": "world"})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsHeteroMap(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"int": types.NewIntValue(-123),
		"str": types.NewStringValue([]byte("hoge")),
		"object": types.NewObjectValue(map[string]*types.Value{
			"nested": types.NewBoolValue(true),
		}),
		"array": types.NewArrayValue([]*types.Value{
			types.NewStringValue([]byte("nested")),
		}),
	})
	got, err := types.ToValue(map[string]interface{}{
		"int": int(-123),
		"str": "hoge",
		"object": map[string]bool{
			"nested": true,
		},
		"array": []interface{}{"nested"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsMapWhoseKeyCanBeConvertedToString(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	got, err := types.ToValue(map[interface{}]string{"hello": "world"})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsSlice(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got, err := types.ToValue([]int{123, 456, 789})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsHeteroSlice(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("hoge")),
		types.NewBoolValue(false),
		types.NewObjectValue(map[string]*types.Value{
			"fuga": types.NewStringValue([]byte("foo")),
		}),
		types.NewArrayValue([]*types.Value{types.NewStringValue([]byte("bar"))}),
	})
	got, err := types.ToValue([]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsArray(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got, err := types.ToValue([3]int{123, 456, 789})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsHeteroArray(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("hoge")),
		types.NewBoolValue(false),
		types.NewObjectValue(map[string]*types.Value{
			"fuga": types.NewStringValue([]byte("foo")),
		}),
		types.NewArrayValue([]*types.Value{types.NewStringValue([]byte("bar"))}),
	})
	got, err := types.ToValue([5]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsPtr(t *testing.T) {
	var i int = 123
	want := types.NewIntValue(123)
	got, err := types.ToValue(&i)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUntaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"name":     types.NewStringValue([]byte("hoge")),
		"longname": types.NewStringValue([]byte("longhoge")),
	})
	got, err := types.ToValue(&untagged{
		Name:     "hoge",
		LongName: "longhoge",
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNestedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"value": types.NewObjectValue(map[string]*types.Value{
			"value": types.NewIntValue(123),
		}),
	})
	got, err := types.ToValue(&nested{
		Value: &nestedInner{
			Value: 123,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsEmbeddedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"field": types.NewIntValue(123),
		"embeddedinner": types.NewObjectValue(map[string]*types.Value{
			"anotherfield": types.NewIntValue(456),
		}),
	})
	value := &embedded{
		Field: 123,
	}
	value.AnotherField = 456
	got, err := types.ToValue(value)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsTaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"customName": types.NewIntValue(123),
	})
	got, err := types.ToValue(&tagged{
		Field: 123,
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueOmitsPrivateField(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"publicfield": types.NewIntValue(123),
	})
	got, err := types.ToValue(&private{
		PublicField:  123,
		privateField: 456,
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueOmitsEmptyFieldTaggedWithOmitempty(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"field1": types.NewIntValue(123),
	})
	f1 := 123
	got, err := types.ToValue(&omitempty{
		Field1: &f1,
		Field2: nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueOmitsFieldTaggedWithHyphen(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"shouldBeIncluded": types.NewIntValue(123),
	})
	got, err := types.ToValue(&alwaysomit{
		ShouldBeIncluded: 123,
		ShouldBeOmitted:  456,
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueEnbedsFieldTaggedWithInline(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"field":       types.NewIntValue(123),
		"nestedfield": types.NewIntValue(456),
	})
	got, err := types.ToValue(&inline{
		Field: 123,
		Inner: inlineInner{
			NestedField: 456,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueUsesMarshalWatsonWhenArgImplementsMarshaler(t *testing.T) {
	m := &customMarshaler{
		SomeField: 123,
	}
	want, err := m.MarshalWatson()
	if err != nil {
		t.Fatal(err)
	}
	got, err := types.ToValue(m)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueUsesMarshalWatsonWhenArgImplementsMarshalerEvenIfNotStruct(t *testing.T) {
	m := primitiveMarshaler(123)
	want, err := m.MarshalWatson()
	if err != nil {
		t.Fatal(err)
	}
	got, err := types.ToValue(m)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilPointer(t *testing.T) {
	var p *int = nil
	want := types.NewNilValue()
	got, err := types.ToValueByReflection(reflect.ValueOf(p))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilSlice(t *testing.T) {
	var p []int = nil
	want := types.NewNilValue()
	got, err := types.ToValueByReflection(reflect.ValueOf(p))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilMap(t *testing.T) {
	var p map[string]interface{} = nil
	want := types.NewNilValue()
	got, err := types.ToValueByReflection(reflect.ValueOf(p))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsTrue(t *testing.T) {
	want := types.NewBoolValue(true)
	got, err := types.ToValueByReflection(reflect.ValueOf(true))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsFalse(t *testing.T) {
	want := types.NewBoolValue(false)
	got, err := types.ToValueByReflection(reflect.ValueOf(false))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt(t *testing.T) {
	want := types.NewIntValue(-12345678)
	got, err := types.ToValueByReflection(reflect.ValueOf(int(-12345678)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt8(t *testing.T) {
	want := types.NewIntValue(64)
	got, err := types.ToValueByReflection(reflect.ValueOf(int8(64)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt16(t *testing.T) {
	want := types.NewIntValue(256)
	got, err := types.ToValueByReflection(reflect.ValueOf(int16(256)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt32(t *testing.T) {
	want := types.NewIntValue(65536)
	got, err := types.ToValueByReflection(reflect.ValueOf(int32(65536)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt64(t *testing.T) {
	want := types.NewIntValue(1234567)
	got, err := types.ToValueByReflection(reflect.ValueOf(int64(1234567)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValueByReflection(reflect.ValueOf(uint(12345)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint8(t *testing.T) {
	want := types.NewUintValue(255)
	got, err := types.ToValueByReflection(reflect.ValueOf(uint8(255)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint16(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValueByReflection(reflect.ValueOf(uint16(12345)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint32(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValueByReflection(reflect.ValueOf(uint32(12345)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint64(t *testing.T) {
	want := types.NewUintValue(12345)
	got, err := types.ToValueByReflection(reflect.ValueOf(uint64(12345)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsString(t *testing.T) {
	want := types.NewStringValue([]byte("hogefuga"))
	got, err := types.ToValueByReflection(reflect.ValueOf("hogefuga"))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsFloat32(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got, err := types.ToValueByReflection(reflect.ValueOf(float32(1.2345e-6)))
	if err != nil {
		t.Fatal(err)
	}
	if got.Kind != types.Float {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueByReflectionConvertsFloat64(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got, err := types.ToValueByReflection(reflect.ValueOf(float64(1.2345e-6)))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsMap(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(map[string]string{"hello": "world"}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsHeteroMap(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"int": types.NewIntValue(-123),
		"str": types.NewStringValue([]byte("hoge")),
		"object": types.NewObjectValue(map[string]*types.Value{
			"nested": types.NewBoolValue(true),
		}),
		"array": types.NewArrayValue([]*types.Value{
			types.NewStringValue([]byte("nested")),
		}),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(map[string]interface{}{
		"int": int(-123),
		"str": "hoge",
		"object": map[string]bool{
			"nested": true,
		},
		"array": []interface{}{"nested"},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsMapWhoseKeyCanBeConvertedToString(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(map[interface{}]string{"hello": "world"}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsSlice(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf([]int{123, 456, 789}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsHeteroSlice(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("hoge")),
		types.NewBoolValue(false),
		types.NewObjectValue(map[string]*types.Value{
			"fuga": types.NewStringValue([]byte("foo")),
		}),
		types.NewArrayValue([]*types.Value{types.NewStringValue([]byte("bar"))}),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf([]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsArray(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf([3]int{123, 456, 789}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsHeteroArray(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123),
		types.NewStringValue([]byte("hoge")),
		types.NewBoolValue(false),
		types.NewObjectValue(map[string]*types.Value{
			"fuga": types.NewStringValue([]byte("foo")),
		}),
		types.NewArrayValue([]*types.Value{types.NewStringValue([]byte("bar"))}),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf([5]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsPtr(t *testing.T) {
	var i int = 123
	want := types.NewIntValue(123)
	got, err := types.ToValueByReflection(reflect.ValueOf(&i))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUntaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"name":     types.NewStringValue([]byte("hoge")),
		"longname": types.NewStringValue([]byte("longhoge")),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(&untagged{
		Name:     "hoge",
		LongName: "longhoge",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNestedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"value": types.NewObjectValue(map[string]*types.Value{
			"value": types.NewIntValue(123),
		}),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(&nested{
		Value: &nestedInner{
			Value: 123,
		},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsEmbeddedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"field": types.NewIntValue(123),
		"embeddedinner": types.NewObjectValue(map[string]*types.Value{
			"anotherfield": types.NewIntValue(456),
		}),
	})
	value := &embedded{
		Field: 123,
	}
	value.AnotherField = 456
	got, err := types.ToValueByReflection(reflect.ValueOf(value))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsTaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"customName": types.NewIntValue(123),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(&tagged{
		Field: 123,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionOmitsPrivateField(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"publicfield": types.NewIntValue(123),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(&private{
		PublicField:  123,
		privateField: 456,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionOmitsEmptyFieldTaggedWithOmitempty(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"field1": types.NewIntValue(123),
	})
	f1 := 123
	got, err := types.ToValueByReflection(reflect.ValueOf(&omitempty{
		Field1: &f1,
		Field2: nil,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionOmitsFieldTaggedWithHyphen(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"shouldBeIncluded": types.NewIntValue(123),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(&alwaysomit{
		ShouldBeIncluded: 123,
		ShouldBeOmitted:  456,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionEnbedsFieldTaggedWithInline(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"field":       types.NewIntValue(123),
		"nestedfield": types.NewIntValue(456),
	})
	got, err := types.ToValueByReflection(reflect.ValueOf(&inline{
		Field: 123,
		Inner: inlineInner{
			NestedField: 456,
		},
	}))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionUsesMarshalWatsonWhenArgImplementsMarshaler(t *testing.T) {
	m := &customMarshaler{
		SomeField: 123,
	}
	want, err := m.MarshalWatson()
	if err != nil {
		t.Fatal(err)
	}
	got, err := types.ToValueByReflection(reflect.ValueOf(m))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByRelectionUsesMarshalWatsonWhenArgImplementsMarshalerEvenIfNotStruct(t *testing.T) {
	m := primitiveMarshaler(123)
	want, err := m.MarshalWatson()
	if err != nil {
		t.Fatal(err)
	}
	got, err := types.ToValueByReflection(reflect.ValueOf(m))
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y)/math.Abs(x) < 1e-3
}
