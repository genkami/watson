package dumper

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/types"
	"github.com/genkami/watson/pkg/vm"
)

func TestDumpInt(t *testing.T) {
	test := func(n int64) {
		orig := types.NewIntValue(n)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test(0)
	test(1)
	test(2)
	test(0x1234abcd)
	test(0x12345678abcdef0)
	test(-1)
}

func TestDumpUint(t *testing.T) {
	test := func(n uint64) {
		orig := types.NewUintValue(n)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test(0)
	test(1)
	test(2)
	test(0x1234abcd)
	test(0x12345678abcdef01)
	test(0xffffffffffffffff)
}

func TestDumpFloat(t *testing.T) {
	test := func(n float64) {
		orig := types.NewFloatValue(n)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if orig.IsNaN() {
			fmt.Printf("yey")
			if !converted.IsNaN() {
				t.Errorf("expected NaN but got %#v", converted)
			}
			return
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test(0)
	test(1)
	test(2)
	test(1.2345e67)
	test(1.2345e-67)
	test(-1.2345e67)
	test(-1.2345e-67)
	test(math.NaN())
	test(math.Inf(1))
	test(math.Inf(-1))
}

func TestDumpString(t *testing.T) {
	test := func(s string) {
		orig := types.NewStringValue([]byte(s))
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test("")
	test("a")
	test("shrimp")
}

func TestDumpObject(t *testing.T) {
	test := func(v map[string]*types.Value) {
		orig := types.NewObjectValue(v)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test(map[string]*types.Value{})
	test(map[string]*types.Value{
		"hoge": types.NewStringValue([]byte("fuga")),
	})
	test(map[string]*types.Value{
		"hoge": types.NewStringValue([]byte("fuga")),
		"fuga": types.NewObjectValue(map[string]*types.Value{
			"foo": types.NewIntValue(0xdeadbeef),
		}),
	})
}

func TestDumpArray(t *testing.T) {
	test := func(arr []*types.Value) {
		orig := types.NewArrayValue(arr)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test([]*types.Value{})
	test([]*types.Value{types.NewIntValue(1)})
	test([]*types.Value{types.NewIntValue(1), types.NewStringValue([]byte("hoge"))})
}

func TestDumpBool(t *testing.T) {
	test := func(b bool) {
		orig := types.NewBoolValue(b)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test(true)
	test(false)
}

func TestDumpNil(t *testing.T) {
	orig := types.NewNilValue()
	converted, err := encodeThenExecute(orig)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(orig, converted); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func encodeThenExecute(val *types.Value) (*types.Value, error) {
	w := lexer.NewSliceWriter()
	d := NewDumper(w)
	err := d.Dump(val)
	if err != nil {
		return nil, err
	}
	dumped := w.Ops()
	v := vm.NewVM()
	for _, op := range dumped {
		err = v.Feed(op)
		if err != nil {
			return nil, err
		}
	}
	return v.Top()
}
