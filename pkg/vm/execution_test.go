package vm

import (
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
