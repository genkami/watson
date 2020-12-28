// Package prettifier decorates Watson Representation by adding some meaningless `vm.Op`s.
package prettifier

import (
	"fmt"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

// Preffifier behaves as a lexer.OpWriter and writes some meaningless Ops to its underlying OpWriter in addition to any Ops that are written.
type Prettifier struct {
	w    lexer.OpWriter
	last *vm.Op
}

// NewPrettifier returns a new Prettifier.
func NewPrettifier(w lexer.OpWriter) *Prettifier {
	return &Prettifier{w: w, last: nil}
}

// Write writes op to the underlying OpWriter.
// It sometimes writes one or more extra Ops to decorate output.
func (p *Prettifier) Write(op vm.Op) error {
	var err error
	if p.last == nil {
		err = p.w.Write(op)
		if err != nil {
			return err
		}
	} else {
		err = p.writeWithDecoration(op, *p.last)
		if err != nil {
			return err
		}
	}
	p.last = &op
	return nil
}

func (p *Prettifier) writeWithDecoration(op vm.Op, last vm.Op) error {
	switch p.Mode() {
	case lexer.A:
		return p.writeWithDecorationA(op, last)
	case lexer.S:
		return p.writeWithDecorationS(op, last)
	default:
		panic(fmt.Errorf("unknown mode: %d", p.Mode()))
	}
}

func (p *Prettifier) writeWithDecorationA(op vm.Op, last vm.Op) error {
	if last == vm.Bnew && op == vm.Oadd {
		return p.writeMulti(vm.Bneg, vm.Bneg, vm.Oadd)
	} else if topShouldBeInt(last) && op == vm.Oadd {
		return p.writeMulti(vm.Ineg, vm.Ineg, vm.Oadd, vm.Gdup, vm.Gpop)
	} else {
		return p.w.Write(op)
	}
}

func (p *Prettifier) writeWithDecorationS(op vm.Op, last vm.Op) error {
	if last == vm.Ishl && op == vm.Iadd { // Sharrk
		return p.writeMulti(vm.Ineg, vm.Ineg, vm.Iadd)
	} else if last == vm.Isht && op == vm.Iadd { // ShaArrk
		return p.writeMulti(vm.Ineg, vm.Ineg, vm.Iadd)
	} else if op == vm.Onew { // Samee+
		return p.writeMulti(vm.Inew, vm.Ishl, vm.Finf, vm.Gpop, vm.Gpop, vm.Onew)
	} else {
		return p.w.Write(op)
	}
}

func (p *Prettifier) writeMulti(ops ...vm.Op) error {
	for _, op := range ops {
		err := p.w.Write(op)
		if err != nil {
			return err
		}
	}
	return nil
}

// Mode returns the Prettifier's current mode.
func (p *Prettifier) Mode() lexer.Mode {
	return p.w.Mode()
}

func topShouldBeInt(op vm.Op) bool {
	for _, v := range []vm.Op{vm.Inew, vm.Iinc, vm.Ishl, vm.Iadd, vm.Ineg, vm.Isht} {
		if op == v {
			return true
		}
	}
	return false
}
