package util

import (
	"testing"

	"github.com/genkami/watson/pkg/vm"

	"github.com/google/go-cmp/cmp"
)

func TestToObjectConvertsInt(t *testing.T) {
	val := vm.NewIntValue(123)
	var want interface{} = int64(123)
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToObjectConvertsFloat(t *testing.T) {
	val := vm.NewFloatValue(1.23)
	var want interface{} = float64(1.23)
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToObjectConvertsString(t *testing.T) {
	val := vm.NewStringValue([]byte("hey"))
	var want interface{} = "hey"
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToObjectConvertsObject(t *testing.T) {
	val := vm.NewObjectValue(map[string]*vm.Value{
		"name": vm.NewStringValue([]byte("Taro")),
		"age":  vm.NewIntValue(25),
	})
	var want interface{} = map[string]interface{}{
		"name": "Taro",
		"age":  int64(25),
	}
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToObjectConvertsArray(t *testing.T) {
	val := vm.NewArrayValue([]*vm.Value{
		vm.NewStringValue([]byte("Yo")),
		vm.NewIntValue(123),
	})
	var want interface{} = []interface{}{"Yo", int64(123)}
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToObjectConvertsBool(t *testing.T) {
	val := vm.NewBoolValue(true)
	var want interface{} = true
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestToObjectConvertsNil(t *testing.T) {
	val := vm.NewNilValue()
	var want interface{} = nil
	got := ToObject(val)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
