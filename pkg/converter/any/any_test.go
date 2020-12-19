package any

import (
	"math"
	"testing"

	"github.com/genkami/watson/pkg/vm"

	"github.com/google/go-cmp/cmp"
)

func TestFromValueConvertsInt(t *testing.T) {
	val := vm.NewIntValue(123)
	var want interface{} = int64(123)
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsFloat(t *testing.T) {
	val := vm.NewFloatValue(1.23)
	var want interface{} = float64(1.23)
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsString(t *testing.T) {
	val := vm.NewStringValue([]byte("hey"))
	var want interface{} = "hey"
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsObject(t *testing.T) {
	val := vm.NewObjectValue(map[string]*vm.Value{
		"name": vm.NewStringValue([]byte("Taro")),
		"age":  vm.NewIntValue(25),
	})
	var want interface{} = map[string]interface{}{
		"name": "Taro",
		"age":  int64(25),
	}
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsArray(t *testing.T) {
	val := vm.NewArrayValue([]*vm.Value{
		vm.NewStringValue([]byte("Yo")),
		vm.NewIntValue(123),
	})
	var want interface{} = []interface{}{"Yo", int64(123)}
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsBool(t *testing.T) {
	val := vm.NewBoolValue(true)
	var want interface{} = true
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsNil(t *testing.T) {
	val := vm.NewNilValue()
	var want interface{} = nil
	got := FromValue(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNil(t *testing.T) {
	want := vm.NewNilValue()
	got := ToValue(nil)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsTrue(t *testing.T) {
	want := vm.NewBoolValue(true)
	got := ToValue(true)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFalse(t *testing.T) {
	want := vm.NewBoolValue(false)
	got := ToValue(false)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt(t *testing.T) {
	want := vm.NewIntValue(-12345678)
	got := ToValue(int(-12345678))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt8(t *testing.T) {
	want := vm.NewIntValue(64)
	got := ToValue(int8(64))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt16(t *testing.T) {
	want := vm.NewIntValue(256)
	got := ToValue(int16(256))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt32(t *testing.T) {
	want := vm.NewIntValue(65536)
	got := ToValue(int32(65536))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt64(t *testing.T) {
	want := vm.NewIntValue(1234567)
	got := ToValue(int64(1234567))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint(t *testing.T) {
	want := vm.NewUintValue(12345)
	got := ToValue(uint(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint8(t *testing.T) {
	want := vm.NewUintValue(255)
	got := ToValue(uint8(255))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint16(t *testing.T) {
	want := vm.NewUintValue(12345)
	got := ToValue(uint16(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint32(t *testing.T) {
	want := vm.NewUintValue(12345)
	got := ToValue(uint32(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint64(t *testing.T) {
	want := vm.NewUintValue(12345)
	got := ToValue(uint64(12345))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsByteSlice(t *testing.T) {
	want := vm.NewStringValue([]byte("hogefuga"))
	got := ToValue([]byte("hogefuga"))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsString(t *testing.T) {
	want := vm.NewStringValue([]byte("hogefuga"))
	got := ToValue("hogefuga")
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFloat32(t *testing.T) {
	want := vm.NewFloatValue(1.2345e-6)
	got := ToValue(float32(1.2345e-6))
	if got.Kind != vm.KFloat {
		t.Fatalf("expected Float but got %#v", got)
	}
	if !closeEnough(want.Float, got.Float) {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestToValueConvertsFloat64(t *testing.T) {
	want := vm.NewFloatValue(1.2345e-6)
	got := ToValue(float64(1.2345e-6))
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsIntMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewIntValue(1),
	})
	got := ToValue(map[string]int{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt8Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewIntValue(1),
	})
	got := ToValue(map[string]int8{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt16Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewIntValue(1),
	})
	got := ToValue(map[string]int16{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt32Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewIntValue(1),
	})
	got := ToValue(map[string]int32{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsInt64Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewIntValue(1),
	})
	got := ToValue(map[string]int64{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUintMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewUintValue(1),
	})
	got := ToValue(map[string]uint{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint8Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewUintValue(1),
	})
	got := ToValue(map[string]uint8{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint16Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewUintValue(1),
	})
	got := ToValue(map[string]uint16{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint32Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewUintValue(1),
	})
	got := ToValue(map[string]uint32{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsUint64Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewUintValue(1),
	})
	got := ToValue(map[string]uint64{
		"value": 1,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsFloat32Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewFloatValue(1.23e-4),
	})
	got := ToValue(map[string]float32{
		"value": 1.23e-4,
	})
	if got.Kind != vm.KObject {
		t.Fatalf("expected Object but got %#v", got)
	}
	gotValue, ok := got.Object["value"]
	if !ok {
		t.Fatalf("missing key: %#v", got)
	}
	if gotValue.Kind != vm.KFloat {
		t.Fatalf("expected float but got: %#v", gotValue)
	}
	if !closeEnough(want.Object["value"].Float, gotValue.Float) {
		t.Errorf("expected %#v but got %#v", want, got)
	}
}

func TestToValueConvertsFloat64Map(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewFloatValue(1.23e-4),
	})
	got := ToValue(map[string]float64{
		"value": 1.23e-4,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilInterfaceMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewNilValue(),
	})
	got := ToValue(map[string]interface{}{
		"value": nil,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

type testStruct struct{}

func TestToValueConvertsNilStructMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewNilValue(),
	})
	got := ToValue(map[string]*testStruct{
		"value": nil,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilMapMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewNilValue(),
	})
	got := ToValue(map[string]map[string]interface{}{
		"value": nil,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsNilSliceMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewNilValue(),
	})
	got := ToValue(map[string][]int{
		"value": nil,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsBoolMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewBoolValue(true),
	})
	got := ToValue(map[string]bool{
		"value": true,
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsStringMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewStringValue([]byte("hoge")),
	})
	got := ToValue(map[string]string{
		"value": "hoge",
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToValueConvertsBytesMap(t *testing.T) {
	want := vm.NewObjectValue(map[string]*vm.Value{
		"value": vm.NewStringValue([]byte("hoge")),
	})
	got := ToValue(map[string][]byte{
		"value": []byte("hoge"),
	})
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func closeEnough(x, y float64) bool {
	return math.Abs(x-y)/math.Abs(x) < 1e-3
}
