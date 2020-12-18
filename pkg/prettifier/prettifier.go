// Package prettifier decorates Watson Representation by adding some meaningless `vm.Op`s.
package prettifier

import (
	"fmt"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

type Prettifier struct {
	w    lexer.OpWriter
	last *vm.Op
}

func NewPrettifier(w lexer.OpWriter) *Prettifier {
	return &Prettifier{w: w, last: nil}
}

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
	return p.w.Write(op)
}

func (p *Prettifier) writeWithDecorationS(op vm.Op, last vm.Op) error {
	if last == vm.Ishl && op == vm.Iadd { // Sharrk
		return p.writeMulti(vm.Ineg, vm.Ineg, vm.Iadd)
	} else if last == vm.Isht && op == vm.Iadd { // ShaArrk
		return p.writeMulti(vm.Ineg, vm.Ineg, vm.Iadd)
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

func (p *Prettifier) Mode() lexer.Mode {
	return p.w.Mode()
}
