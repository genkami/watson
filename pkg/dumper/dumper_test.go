package dumper

import (
	"testing"

	"github.com/genkami/watson/pkg/vm"
	"github.com/google/go-cmp/cmp"
)

func TestSliceWritersInitialOpsIsEmpty(t *testing.T) {
	w := NewSliceWriter()
	ops := w.Ops()
	if len(ops) != 0 {
		t.Errorf("expected empty slice but got %#v", ops)
	}
}

func TestSliceWriterReturnsAllOpsThatAreWritten(t *testing.T) {
	w := NewSliceWriter()
	expected := []vm.Op{vm.Inew, vm.Iinc, vm.Ineg, vm.Fneg, vm.Snew, vm.Sadd}
	for _, op := range expected {
		err := w.Write(op)
		if err != nil {
			t.Fatal(err)
		}
	}
	actual := w.Ops()
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
