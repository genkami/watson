// Package dumper converts `vm.Value`s into `vm.Op`s.
package dumper

import (
	"fmt"
	"math"

	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

// Dumper dumps `vm.Value` as a sequence of `vm.Op`s.
type Dumper struct {
	w lexer.OpWriter
}

// NewDumper creates a new Dumper.
func NewDumper(w lexer.OpWriter) *Dumper {
	return &Dumper{w: w}
}

// Dump converts v into a sequence of `vm.Op`s and writes it to the underlying writer `lexer.OpWriter`.
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
	case vm.KArray:
		return d.dumpArray(v.Array)
	case vm.KBool:
		return d.dumpBool(v.Bool)
	case vm.KNil:
		return d.dumpNil()
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

func (d *Dumper) dumpArray(arr []*vm.Value) error {
	var err error
	err = d.w.Write(vm.Anew)
	if err != nil {
		return err
	}
	for _, v := range arr {
		err = d.Dump(v)
		if err != nil {
			return err
		}
		err = d.w.Write(vm.Aadd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dumper) dumpBool(b bool) error {
	var err error
	err = d.w.Write(vm.Bnew)
	if err != nil {
		return err
	}
	if b {
		err = d.w.Write(vm.Bneg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dumper) dumpNil() error {
	return d.w.Write(vm.Nnew)
}
