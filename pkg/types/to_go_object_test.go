package types_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/types"
)

func TestFromValueConvertsInt(t *testing.T) {
	val := types.NewIntValue(123)
	var want interface{} = int64(123)
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsFloat(t *testing.T) {
	val := types.NewFloatValue(1.23)
	var want interface{} = float64(1.23)
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsString(t *testing.T) {
	val := types.NewStringValue([]byte("hey"))
	var want interface{} = "hey"
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsObject(t *testing.T) {
	val := types.NewObjectValue(map[string]*types.Value{
		"name": types.NewStringValue([]byte("Taro")),
		"age":  types.NewIntValue(25),
	})
	var want interface{} = map[string]interface{}{
		"name": "Taro",
		"age":  int64(25),
	}
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsArray(t *testing.T) {
	val := types.NewArrayValue([]*types.Value{
		types.NewStringValue([]byte("Yo")),
		types.NewIntValue(123),
	})
	var want interface{} = []interface{}{"Yo", int64(123)}
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsBool(t *testing.T) {
	val := types.NewBoolValue(true)
	var want interface{} = true
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFromValueConvertsNil(t *testing.T) {
	val := types.NewNilValue()
	var want interface{} = nil
	got := val.ToGoObject()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
