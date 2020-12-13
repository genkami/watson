package lexer

import (
	"bytes"
	"io"
	"testing"

	"github.com/genkami/watson/pkg/vm"

	"github.com/google/go-cmp/cmp"
)

func TestNewLexer(t *testing.T) {
	lex := NewLexer(bytes.NewReader(nil))
	if lex.Mode() != A {
		t.Fatalf("expected %#v but got %#v", A, lex.Mode())
	}
}

func TestNewLexerWithInitialModeSetToA(t *testing.T) {
	lex := NewLexer(bytes.NewReader(nil), WithInitialMode(A))
	if lex.Mode() != A {
		t.Fatalf("expected %#v but got %#v", A, lex.Mode())
	}
}

func TestNewLexerWithInitialModeSetToS(t *testing.T) {
	lex := NewLexer(bytes.NewReader(nil), WithInitialMode(S))
	if lex.Mode() != S {
		t.Fatalf("expected %#v but got %#v", S, lex.Mode())
	}
}

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

func TestOpTableIsSurjectiveWhenModeIsS(t *testing.T) {
	ops := map[vm.Op]bool{}
	for _, op := range vm.AllOps() {
		ops[op] = true
	}
	for _, op := range opTableS {
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

func TestNextChangesItsStateFromAToCWhenReachingSnew(t *testing.T) {
	got, err := readAll("b?b")
	if err != nil {
		t.Fatal(err)
	}
	want := []vm.Op{vm.Ishl, vm.Snew, vm.Fnan}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func TestNextChangesItsStateFromCToAWhenTheNextTimeItReachesSnew(t *testing.T) {
	got, err := readAll("b?b$b")
	if err != nil {
		t.Fatal(err)
	}
	want := []vm.Op{vm.Ishl, vm.Snew, vm.Fnan, vm.Snew, vm.Ishl}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("expected %#v but got %#v", want, got)
	}
}

func readOne(s string) (vm.Op, error) {
	buf := bytes.NewReader([]byte(s))
	l := NewLexer(buf)
	return l.Next()
}

func readAll(s string) ([]vm.Op, error) {
	buf := bytes.NewReader([]byte(s))
	l := NewLexer(buf)
	out := make([]vm.Op, 0, 10) // a random number that is sufficient to run the test
	for {
		op, err := l.Next()
		if err != nil {
			if err == io.EOF {
				return out, nil
			} else {
				return nil, err
			}
		}
		out = append(out, op)
	}
}
