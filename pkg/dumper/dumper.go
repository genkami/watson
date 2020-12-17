// Package dumper converts `vm.Value`s into `vm.Op`s.
package dumper

import (
	"github.com/genkami/watson/pkg/vm"
)

// OpWriter is an underlying writer that is used by Dumper.
// In most cases it would be `lexer.Unlexer`.
type OpWriter interface {
	Write(vm.Op) error
}

// SliceWriter is a simple `OpWriter` that just holds `vm.Op`s written as a slice of `vm.Op`s.
type SliceWriter struct {
	ops []vm.Op
}

// NewSliceWriter creates a new `SliceWriter`.
func NewSliceWriter() *SliceWriter {
	return &SliceWriter{ops: make([]vm.Op, 0, 50)}
}

// Returns the `vm.Op`s that was written by previous `Write`s.
func (s *SliceWriter) Ops() []vm.Op {
	ops := make([]vm.Op, len(s.ops))
	copy(ops, s.ops)
	return ops
}

func (s *SliceWriter) Write(op vm.Op) error {
	s.ops = append(s.ops, op)
	return nil
}

// Dumper dumps `vm.Value` as a sequence of `vm.Op`s.
type Dumper struct {
	w OpWriter
}

// NewDumper creates a new Dumper.
func NewDumper(w OpWriter) *Dumper {
	return &Dumper{w: w}
}

// Dump converts v into a sequence of `vm.Op`s and writes it to the underlying writer `OpWriter`.
func (d *Dumper) Dump(v *vm.Value) error {
	return nil
}
