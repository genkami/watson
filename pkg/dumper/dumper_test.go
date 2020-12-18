package dumper

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

func TestDumpInt(t *testing.T) {
	test := func(n int64) {
		orig := vm.NewIntValue(n)
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
		orig := vm.NewUintValue(n)
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
		orig := vm.NewFloatValue(n)
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
		orig := vm.NewStringValue([]byte(s))
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
	test := func(v map[string]*vm.Value) {
		orig := vm.NewObjectValue(v)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test(map[string]*vm.Value{})
	test(map[string]*vm.Value{
		"hoge": vm.NewStringValue([]byte("fuga")),
	})
	test(map[string]*vm.Value{
		"hoge": vm.NewStringValue([]byte("fuga")),
		"fuga": vm.NewObjectValue(map[string]*vm.Value{
			"foo": vm.NewIntValue(0xdeadbeef),
		}),
	})
}

func TestDumpArray(t *testing.T) {
	test := func(arr []*vm.Value) {
		orig := vm.NewArrayValue(arr)
		converted, err := encodeThenExecute(orig)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(orig, converted); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
	test([]*vm.Value{})
	test([]*vm.Value{vm.NewIntValue(1)})
	test([]*vm.Value{vm.NewIntValue(1), vm.NewStringValue([]byte("hoge"))})
}

func TestDumpBool(t *testing.T) {
	test := func(b bool) {
		orig := vm.NewBoolValue(b)
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
	orig := vm.NewNilValue()
	converted, err := encodeThenExecute(orig)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(orig, converted); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func encodeThenExecute(val *vm.Value) (*vm.Value, error) {
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
