package lexer

import (
	"bytes"
	"io"
	"testing"

	"github.com/genkami/watson/pkg/vm"
)

func TestOpTableIsSurjectiveWhenModeIsA(t *testing.T) {
	ops := map[vm.Op]bool{}
	for _, op := range vm.AllOps() {
		ops[op] = true
	}
	for _, op := range opTableA {
		delete(ops, op)
	}
	for op := range ops {
		t.Errorf("%#v is not in opTable", op)
	}
}

func TestNextReturnsTheFirstOp(t *testing.T) {
	op, err := readOne("B")
	if err != nil {
		t.Fatal(err)
	}
	if op != vm.Inew {
		t.Errorf("expected %#v but got %#v", vm.Inew, op)
	}
}

func TestNextReturnsOpsSequentially(t *testing.T) {
	buf := bytes.NewReader([]byte("Bubba"))
	l := NewLexer(buf)
	expectedOps := []vm.Op{vm.Inew, vm.Iinc, vm.Ishl, vm.Ishl, vm.Iadd}
	for _, expected := range expectedOps {
		actual, err := l.Next()
		if err != nil {
			t.Fatal(err)
		}
		if expected != actual {
			t.Errorf("expected %#v but got %#v", expected, actual)
		}
	}
	_, err := l.Next()
	if err != io.EOF {
		t.Fatal(err)
	}
}

func TestNextSkipsMeaninglessBytes(t *testing.T) {
	op, err := readOne("ZZZZZB")
	if err != nil {
		t.Fatal(err)
	}
	if op != vm.Inew {
		t.Errorf("expected %#v but got %#v", vm.Inew, op)
	}
}

func TestNextReturnsEOFWhenReaderIsEmpty(t *testing.T) {
	_, err := readOne("")
	if err != io.EOF {
		t.Fatal(err)
	}
}

func TestNextReturnsEOFWhenReachingEndOfFile(t *testing.T) {
	_, err := readOne("ZZZZZZZZ")
	if err != io.EOF {
		t.Fatal(err)
	}
}

func readOne(s string) (vm.Op, error) {
	buf := bytes.NewReader([]byte(s))
	l := NewLexer(buf)
	return l.Next()
}
