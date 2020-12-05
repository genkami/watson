package vm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFeedInewPushesZero(t *testing.T) {
	vm := NewVM()
	err := vm.Feed(Inew)
	if err != nil {
		t.Fatal(err)
	}

	if vm.sp != 0 {
		t.Fatalf("stack pointer mismatch: expected 1, got %d", vm.sp)
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
