package types

import (
	"math"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToValueConvertsNilInterface(t *testing.T) {
	want := NewNilValue()
	got := ToValue(nil)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilPointer(t *testing.T) {
	var p *int = nil
	want := NewNilValue()
	got := ToValue(p)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilSlice(t *testing.T) {
	var p []int = nil
	want := NewNilValue()
	got := ToValue(p)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilMap(t *testing.T) {
	var p map[string]interface{} = nil
	want := NewNilValue()
	got := ToValue(p)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsTrue(t *testing.T) {
	want := NewBoolValue(true)
	got := ToValue(true)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFalse(t *testing.T) {
	want := NewBoolValue(false)
	got := ToValue(false)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt(t *testing.T) {
	want := NewIntValue(-12345678)
	got := ToValue(int(-12345678))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt8(t *testing.T) {
	want := NewIntValue(64)
	got := ToValue(int8(64))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt16(t *testing.T) {
	want := NewIntValue(256)
	got := ToValue(int16(256))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt32(t *testing.T) {
	want := NewIntValue(65536)
	got := ToValue(int32(65536))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt64(t *testing.T) {
	want := NewIntValue(1234567)
	got := ToValue(int64(1234567))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValue(uint(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint8(t *testing.T) {
	want := NewUintValue(255)
	got := ToValue(uint8(255))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint16(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValue(uint16(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint32(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValue(uint32(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint64(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValue(uint64(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsString(t *testing.T) {
	want := NewStringValue([]byte("hogefuga"))
	got := ToValue("hogefuga")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFloat32(t *testing.T) {
	want := NewFloatValue(1.2345e-6)
	got := ToValue(float32(1.2345e-6))
	if got.Kind != Float {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueConvertsFloat64(t *testing.T) {
	want := NewFloatValue(1.2345e-6)
	got := ToValue(float64(1.2345e-6))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsMap(t *testing.T) {
	want := NewObjectValue(map[string]*Value{
		"hello": NewStringValue([]byte("world")),
	})
	got := ToValue(map[string]string{"hello": "world"})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsHeteroMap(t *testing.T) {
	want := NewObjectValue(map[string]*Value{
		"int": NewIntValue(-123),
		"str": NewStringValue([]byte("hoge")),
		"object": NewObjectValue(map[string]*Value{
			"nested": NewBoolValue(true),
		}),
		"array": NewArrayValue([]*Value{
			NewStringValue([]byte("nested")),
		}),
	})
	got := ToValue(map[string]interface{}{
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
	want := NewArrayValue([]*Value{
		NewIntValue(123), NewIntValue(456), NewIntValue(789),
	})
	got := ToValue([]int{123, 456, 789})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsHeteroSlice(t *testing.T) {
	want := NewArrayValue([]*Value{
		NewIntValue(123),
		NewStringValue([]byte("hoge")),
		NewBoolValue(false),
		NewObjectValue(map[string]*Value{
			"fuga": NewStringValue([]byte("foo")),
		}),
		NewArrayValue([]*Value{NewStringValue([]byte("bar"))}),
	})
	got := ToValue([]interface{}{
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
	want := NewArrayValue([]*Value{
		NewIntValue(123), NewIntValue(456), NewIntValue(789),
	})
	got := ToValue([3]int{123, 456, 789})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsHeteroArray(t *testing.T) {
	want := NewArrayValue([]*Value{
		NewIntValue(123),
		NewStringValue([]byte("hoge")),
		NewBoolValue(false),
		NewObjectValue(map[string]*Value{
			"fuga": NewStringValue([]byte("foo")),
		}),
		NewArrayValue([]*Value{NewStringValue([]byte("bar"))}),
	})
	got := ToValue([5]interface{}{
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
	want := NewIntValue(123)
	got := ToValue(&i)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilPointer(t *testing.T) {
	var p *int = nil
	want := NewNilValue()
	got := ToValueByReflection(reflect.ValueOf(p))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilSlice(t *testing.T) {
	var p []int = nil
	want := NewNilValue()
	got := ToValueByReflection(reflect.ValueOf(p))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsNilMap(t *testing.T) {
	var p map[string]interface{} = nil
	want := NewNilValue()
	got := ToValueByReflection(reflect.ValueOf(p))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsTrue(t *testing.T) {
	want := NewBoolValue(true)
	got := ToValueByReflection(reflect.ValueOf(true))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsFalse(t *testing.T) {
	want := NewBoolValue(false)
	got := ToValueByReflection(reflect.ValueOf(false))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt(t *testing.T) {
	want := NewIntValue(-12345678)
	got := ToValueByReflection(reflect.ValueOf(int(-12345678)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt8(t *testing.T) {
	want := NewIntValue(64)
	got := ToValueByReflection(reflect.ValueOf(int8(64)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt16(t *testing.T) {
	want := NewIntValue(256)
	got := ToValueByReflection(reflect.ValueOf(int16(256)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt32(t *testing.T) {
	want := NewIntValue(65536)
	got := ToValueByReflection(reflect.ValueOf(int32(65536)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsInt64(t *testing.T) {
	want := NewIntValue(1234567)
	got := ToValueByReflection(reflect.ValueOf(int64(1234567)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValueByReflection(reflect.ValueOf(uint(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint8(t *testing.T) {
	want := NewUintValue(255)
	got := ToValueByReflection(reflect.ValueOf(uint8(255)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint16(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValueByReflection(reflect.ValueOf(uint16(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint32(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValueByReflection(reflect.ValueOf(uint32(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsUint64(t *testing.T) {
	want := NewUintValue(12345)
	got := ToValueByReflection(reflect.ValueOf(uint64(12345)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsString(t *testing.T) {
	want := NewStringValue([]byte("hogefuga"))
	got := ToValueByReflection(reflect.ValueOf("hogefuga"))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsFloat32(t *testing.T) {
	want := NewFloatValue(1.2345e-6)
	got := ToValueByReflection(reflect.ValueOf(float32(1.2345e-6)))
	if got.Kind != Float {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueByReflectionConvertsFloat64(t *testing.T) {
	want := NewFloatValue(1.2345e-6)
	got := ToValueByReflection(reflect.ValueOf(float64(1.2345e-6)))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsMap(t *testing.T) {
	want := NewObjectValue(map[string]*Value{
		"hello": NewStringValue([]byte("world")),
	})
	got := ToValueByReflection(reflect.ValueOf(map[string]string{"hello": "world"}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsHeteroMap(t *testing.T) {
	want := NewObjectValue(map[string]*Value{
		"int": NewIntValue(-123),
		"str": NewStringValue([]byte("hoge")),
		"object": NewObjectValue(map[string]*Value{
			"nested": NewBoolValue(true),
		}),
		"array": NewArrayValue([]*Value{
			NewStringValue([]byte("nested")),
		}),
	})
	got := ToValueByReflection(reflect.ValueOf(map[string]interface{}{
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
	want := NewArrayValue([]*Value{
		NewIntValue(123), NewIntValue(456), NewIntValue(789),
	})
	got := ToValueByReflection(reflect.ValueOf([]int{123, 456, 789}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsHeteroSlice(t *testing.T) {
	want := NewArrayValue([]*Value{
		NewIntValue(123),
		NewStringValue([]byte("hoge")),
		NewBoolValue(false),
		NewObjectValue(map[string]*Value{
			"fuga": NewStringValue([]byte("foo")),
		}),
		NewArrayValue([]*Value{NewStringValue([]byte("bar"))}),
	})
	got := ToValueByReflection(reflect.ValueOf([]interface{}{
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
	want := NewArrayValue([]*Value{
		NewIntValue(123), NewIntValue(456), NewIntValue(789),
	})
	got := ToValueByReflection(reflect.ValueOf([3]int{123, 456, 789}))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueByReflectionConvertsHeteroArray(t *testing.T) {
	want := NewArrayValue([]*Value{
		NewIntValue(123),
		NewStringValue([]byte("hoge")),
		NewBoolValue(false),
		NewObjectValue(map[string]*Value{
			"fuga": NewStringValue([]byte("foo")),
		}),
		NewArrayValue([]*Value{NewStringValue([]byte("bar"))}),
	})
	got := ToValueByReflection(reflect.ValueOf([5]interface{}{
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
	want := NewIntValue(123)
	got := ToValueByReflection(reflect.ValueOf(&i))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y)/math.Abs(x) < 1e-3
}
