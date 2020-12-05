package vm

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFeedInewPushesZero(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Inew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewIntValue(0)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIincIncrementsTheValue(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Inew)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iinc)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewIntValue(1)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIincFailsWhenTypeMismatch(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Nnew)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iinc)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIincFailsIfStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Iinc)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIshlShiftsTheTopBy1(t *testing.T) {
	var err error
	vm := NewVM()

	var before int64 = 123

	err = vm.pushInt(before)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Ishl)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewIntValue(before * 2)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIshlFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.Feed(Ishl)
	if err != ErrStackEmpty {
		t.Fatal(err)
	}
}

func TestFeedIshlFailsWhenTypeMismatch(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Ishl)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIaddAddsTwoIntegers(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(2)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewIntValue(3)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedIaddFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Iadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIaddFailsWhenStackIsInsufficiient(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedIaddFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedIaddFailsWhenArg2IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Iadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedSnewPushesEmptyString(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Snew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewStringValue([]byte{})
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedSaddAddsACharToString(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushString([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(0x21) // '!'
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewStringValue([]byte("hello!"))
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedSaddFailsWhenStackIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Sadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedSaddFailsWhenStackIsInsufficiient(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.pushInt(1)
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != ErrStackEmpty {
		t.Fatalf("expected ErrStackEmpty but got %v", err)
	}
}

func TestFeedSaddFailsWhenArg1IsNotInteger(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushString([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedSaddFailsWhenArg2IsNotString(t *testing.T) {
	var err error
	vm := NewVM()

	err = vm.pushNil()
	if err != nil {
		t.Fatal(err)
	}
	err = vm.pushInt(0x21) // '!'
	if err != nil {
		t.Fatal(err)
	}
	err = vm.Feed(Sadd)
	if err != ErrTypeMismatch {
		t.Fatal(err)
	}
}

func TestFeedNnewPushesNil(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.Feed(Nnew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", 0, vm.sp)
	}

	want := NewNilValue()
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestFeedMultiDoNothingWhenOpsIsEmpty(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.FeedMulti([]Op{})
	if err != nil {
		t.Fatal(err)
	}
	if vm.sp != -1 {
		t.Fatalf("stack pointer mismatch: expected %d, got %d", -1, vm.sp)
	}
}

func TestFeedMultiExecutesOpsSequentially(t *testing.T) {
	var err error
	vm := NewVM()
	err = vm.FeedMulti([]Op{Inew, Iinc, Iinc, Iinc})
	if err != nil {
		t.Fatal(err)
	}

	want := NewIntValue(3)
	got, err := vm.Top()
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestDeepCopyWithInt(t *testing.T) {
	orig := NewIntValue(123)
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Int = 456
	if orig.Int == clone.Int {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyString(t *testing.T) {
	orig := NewStringValue([]byte("hello"))
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.String[0] = 0x61 // 'a'
	if bytes.Equal(orig.String, clone.String) {
		t.Errorf("clone shares the same reference with its origin")
	}

	clone.String = []byte("world")
	if bytes.Equal(orig.String, clone.String) {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyWithObject(t *testing.T) {
	orig := NewObjectValue(map[string]*Value{
		"hello": NewStringValue([]byte("world")),
	})
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Object["hello"].String[0] = 0x61 // 'a'
	if diff := cmp.Diff(orig, clone); diff == "" {
		t.Errorf("clone shares the same reference with its origin")
	}

	clone.Object["hoge"] = NewStringValue([]byte("fuga"))
	if diff := cmp.Diff(orig, clone); diff == "" {
		t.Errorf("clone shares the same reference with its origin")
	}

	clone.Object = map[string]*Value{}
	if diff := cmp.Diff(orig, clone); diff == "" {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyWithNil(t *testing.T) {
	orig := NewNilValue()
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Kind = KInt
	if orig.Kind == clone.Kind {
		t.Errorf("DeepCopy returned receiver itself")
	}
}
