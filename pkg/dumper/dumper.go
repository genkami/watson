// Package dumper converts `vm.Value`s into `vm.Op`s.
package dumper

import (
	"fmt"
	"math"

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
	switch v.Kind {
	case vm.KInt:
		return d.dumpInt(uint64(v.Int))
	case vm.KFloat:
		return d.dumpFloat(v.Float)
	case vm.KString:
		return d.dumpString(v.String)
	case vm.KObject:
		return d.dumpObject(v.Object)
	default:
		panic(fmt.Errorf("unknown kind: %d", v.Kind))
	}
}

func (d *Dumper) dumpInt(n uint64) error {
	var err error
	err = d.w.Write(vm.Inew)
	if err != nil {
		return err
	}
	shift := 0
	for n != 0 {
		if n%2 == 1 {
			err = d.w.Write(vm.Inew)
			if err != nil {
				return err
			}
			err = d.w.Write(vm.Iinc)
			if err != nil {
				return err
			}
			for i := 0; i < shift; i++ {
				err = d.w.Write(vm.Ishl)
				if err != nil {
					return err
				}
			}
			err = d.w.Write(vm.Iadd)
			if err != nil {
				return err
			}
		}
		n = n >> 1
		shift++
	}
	return nil
}

func (d *Dumper) dumpFloat(x float64) error {
	var err error
	if math.IsNaN(x) {
		return d.w.Write(vm.Fnan)
	} else if math.IsInf(x, 1) {
		return d.w.Write(vm.Finf)
	} else if math.IsInf(x, -1) {
		err = d.w.Write(vm.Finf)
		if err != nil {
			return err
		}
		err = d.w.Write(vm.Fneg)
		if err != nil {
			return err
		}
	}
	err = d.dumpInt(math.Float64bits(x))
	if err != nil {
		return err
	}
	err = d.w.Write(vm.Itof)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dumper) dumpString(s []byte) error {
	var err error
	err = d.w.Write(vm.Snew)
	if err != nil {
		return err
	}
	for _, c := range s {
		err = d.dumpInt(uint64(c))
		if err != nil {
			return err
		}
		err = d.w.Write(vm.Sadd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dumper) dumpObject(obj map[string]*vm.Value) error {
	var err error
	err = d.w.Write(vm.Onew)
	if err != nil {
		return err
	}
	for k, v := range obj {
		err = d.dumpString([]byte(k))
		if err != nil {
			return err
		}
		err = d.Dump(v)
		if err != nil {
			return err
		}
		err = d.w.Write(vm.Oadd)
		if err != nil {
			return err
		}
	}
	return nil
}