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

func TestNewLexerWithFileNameSetsFileName(t *testing.T) {
	name := "hoge.txt"
	lex := NewLexer(bytes.NewReader(nil), WithFileName(name))
	if lex.fileName != name {
		t.Fatalf("expected %#v but got %#v", name, lex.fileName)
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

func TestNextReturnsFileNameAndPosition(t *testing.T) {
	name := "hoge.watson"
	buf := bytes.NewReader([]byte("Bub\nba"))
	l := NewLexer(buf, WithFileName(name))
	expectedTokens := []*Token{
		&Token{Op: vm.Inew, FileName: name, Line: 0, Column: 0},
		&Token{Op: vm.Iinc, FileName: name, Line: 0, Column: 1},
		&Token{Op: vm.Ishl, FileName: name, Line: 0, Column: 2},
		&Token{Op: vm.Ishl, FileName: name, Line: 1, Column: 0},
		&Token{Op: vm.Iadd, FileName: name, Line: 1, Column: 1},
	}
	for _, expected := range expectedTokens {
		actual, err := l.Next()
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(expected, actual); diff != "" {
			t.Errorf("lexing %#v: mismatch (-want +got):\n%s", expected.Op, diff)
		}
	}
}

func TestNextReturnsOpsSequentially(t *testing.T) {
	buf := bytes.NewReader([]byte("Bubba"))
	l := NewLexer(buf)
	expectedOps := []vm.Op{vm.Inew, vm.Iinc, vm.Ishl, vm.Ishl, vm.Iadd}
	for _, expected := range expectedOps {
		tok, err := l.Next()
		if err != nil {
			t.Fatal(err)
		}
		actual := tok.Op
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
	tok, err := l.Next()
	if err != nil {
		return 0, err
	}
	return tok.Op, nil
}

func readAll(s string) ([]vm.Op, error) {
	buf := bytes.NewReader([]byte(s))
	l := NewLexer(buf)
	out := make([]vm.Op, 0, 10) // a random number that is sufficient to run the test
	for {
		tok, err := l.Next()
		if err != nil {
			if err == io.EOF {
				return out, nil
			} else {
				return nil, err
			}
		}
		out = append(out, tok.Op)
	}
}
