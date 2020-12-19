package types_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/types"
)

type untagged struct {
	Name     string
	LongName string
}

type nested struct {
	value *nestedInner
}

type nestedInner struct {
	value int
}

type embedded struct {
	field int
	embeddedInner
}

type embeddedInner struct {
	anotherField int
}

type tagged struct {
	Field int `watson:"customName"`
}

func TestToValueConvertsNilInterface(t *testing.T) {
	want := types.NewNilValue()
	got := types.ToValue(nil)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilPointer(t *testing.T) {
	var p *int = nil
	want := types.NewNilValue()
	got := types.ToValue(p)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilSlice(t *testing.T) {
	var p []int = nil
	want := types.NewNilValue()
	got := types.ToValue(p)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilMap(t *testing.T) {
	var p map[string]interface{} = nil
	want := types.NewNilValue()
	got := types.ToValue(p)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsTrue(t *testing.T) {
	want := types.NewBoolValue(true)
	got := types.ToValue(true)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFalse(t *testing.T) {
	want := types.NewBoolValue(false)
	got := types.ToValue(false)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt(t *testing.T) {
	want := types.NewIntValue(-12345678)
	got := types.ToValue(int(-12345678))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt8(t *testing.T) {
	want := types.NewIntValue(64)
	got := types.ToValue(int8(64))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt16(t *testing.T) {
	want := types.NewIntValue(256)
	got := types.ToValue(int16(256))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt32(t *testing.T) {
	want := types.NewIntValue(65536)
	got := types.ToValue(int32(65536))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt64(t *testing.T) {
	want := types.NewIntValue(1234567)
	got := types.ToValue(int64(1234567))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValue(uint(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint8(t *testing.T) {
	want := types.NewUintValue(255)
	got := types.ToValue(uint8(255))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint16(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValue(uint16(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint32(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValue(uint32(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint64(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValue(uint64(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsString(t *testing.T) {
	want := types.NewStringValue([]byte("hogefuga"))
	got := types.ToValue("hogefuga")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFloat32(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got := types.ToValue(float32(1.2345e-6))
	if got.Kind != types.Float {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueConvertsFloat64(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got := types.ToValue(float64(1.2345e-6))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsMap(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	got := types.ToValue(map[string]string{"hello": "world"})
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
	got := types.ToValue(map[string]interface{}{
		"int": int(-123),
		"str": "hoge",
		"object": map[string]bool{
			"nested": true,
		},
		"array": []interface{}{"nested"},
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsSlice(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got := types.ToValue([]int{123, 456, 789})
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
	got := types.ToValue([]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsArray(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got := types.ToValue([3]int{123, 456, 789})
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
	got := types.ToValue([5]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsPtr(t *testing.T) {
	var i int = 123
	want := types.NewIntValue(123)
	got := types.ToValue(&i)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUntaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"name":     types.NewStringValue([]byte("hoge")),
		"longname": types.NewStringValue([]byte("longhoge")),
	})
	got := types.ToValue(&untagged{
		Name:     "hoge",
		LongName: "longhoge",
	})
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
	got := types.ToValue(&nested{
		value: &nestedInner{
			value: 123,
		},
	})
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
		field: 123,
	}
	value.anotherField = 456
	got := types.ToValue(value)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsTaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"customName": types.NewIntValue(123),
	})
	got := types.ToValue(&tagged{
		Field: 123,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilPointer(t *testing.T) {
	var p *int = nil
	want := types.NewNilValue()
	got := types.ToValueByReflection(reflect.ValueOf(p))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilSlice(t *testing.T) {
	var p []int = nil
	want := types.NewNilValue()
	got := types.ToValueByReflection(reflect.ValueOf(p))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilMap(t *testing.T) {
	var p map[string]interface{} = nil
	want := types.NewNilValue()
	got := types.ToValueByReflection(reflect.ValueOf(p))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsTrue(t *testing.T) {
	want := types.NewBoolValue(true)
	got := types.ToValueByReflection(reflect.ValueOf(true))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsFalse(t *testing.T) {
	want := types.NewBoolValue(false)
	got := types.ToValueByReflection(reflect.ValueOf(false))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt(t *testing.T) {
	want := types.NewIntValue(-12345678)
	got := types.ToValueByReflection(reflect.ValueOf(int(-12345678)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt8(t *testing.T) {
	want := types.NewIntValue(64)
	got := types.ToValueByReflection(reflect.ValueOf(int8(64)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt16(t *testing.T) {
	want := types.NewIntValue(256)
	got := types.ToValueByReflection(reflect.ValueOf(int16(256)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt32(t *testing.T) {
	want := types.NewIntValue(65536)
	got := types.ToValueByReflection(reflect.ValueOf(int32(65536)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt64(t *testing.T) {
	want := types.NewIntValue(1234567)
	got := types.ToValueByReflection(reflect.ValueOf(int64(1234567)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValueByReflection(reflect.ValueOf(uint(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint8(t *testing.T) {
	want := types.NewUintValue(255)
	got := types.ToValueByReflection(reflect.ValueOf(uint8(255)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint16(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValueByReflection(reflect.ValueOf(uint16(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint32(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValueByReflection(reflect.ValueOf(uint32(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint64(t *testing.T) {
	want := types.NewUintValue(12345)
	got := types.ToValueByReflection(reflect.ValueOf(uint64(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsString(t *testing.T) {
	want := types.NewStringValue([]byte("hogefuga"))
	got := types.ToValueByReflection(reflect.ValueOf("hogefuga"))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsFloat32(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got := types.ToValueByReflection(reflect.ValueOf(float32(1.2345e-6)))
	if got.Kind != types.Float {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueByReflectionConvertsFloat64(t *testing.T) {
	want := types.NewFloatValue(1.2345e-6)
	got := types.ToValueByReflection(reflect.ValueOf(float64(1.2345e-6)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsMap(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"hello": types.NewStringValue([]byte("world")),
	})
	got := types.ToValueByReflection(reflect.ValueOf(map[string]string{"hello": "world"}))
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
	got := types.ToValueByReflection(reflect.ValueOf(map[string]interface{}{
		"int": int(-123),
		"str": "hoge",
		"object": map[string]bool{
			"nested": true,
		},
		"array": []interface{}{"nested"},
	}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsSlice(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got := types.ToValueByReflection(reflect.ValueOf([]int{123, 456, 789}))
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
	got := types.ToValueByReflection(reflect.ValueOf([]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsArray(t *testing.T) {
	want := types.NewArrayValue([]*types.Value{
		types.NewIntValue(123), types.NewIntValue(456), types.NewIntValue(789),
	})
	got := types.ToValueByReflection(reflect.ValueOf([3]int{123, 456, 789}))
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
	got := types.ToValueByReflection(reflect.ValueOf([5]interface{}{
		int(123),
		"hoge",
		false,
		map[string]interface{}{
			"fuga": "foo",
		},
		[]string{"bar"},
	}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsPtr(t *testing.T) {
	var i int = 123
	want := types.NewIntValue(123)
	got := types.ToValueByReflection(reflect.ValueOf(&i))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUntaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"name":     types.NewStringValue([]byte("hoge")),
		"longname": types.NewStringValue([]byte("longhoge")),
	})
	got := types.ToValueByReflection(reflect.ValueOf(&untagged{
		Name:     "hoge",
		LongName: "longhoge",
	}))
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
	got := types.ToValueByReflection(reflect.ValueOf(&nested{
		value: &nestedInner{
			value: 123,
		},
	}))
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
		field: 123,
	}
	value.anotherField = 456
	got := types.ToValueByReflection(reflect.ValueOf(value))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsTaggedStruct(t *testing.T) {
	want := types.NewObjectValue(map[string]*types.Value{
		"customName": types.NewIntValue(123),
	})
	got := types.ToValueByReflection(reflect.ValueOf(&tagged{
		Field: 123,
	}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y)/math.Abs(x) < 1e-3
}
