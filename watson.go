package watson

import (
	"io"

	"github.com/genkami/watson/pkg/dumper"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/types"
	"github.com/genkami/watson/pkg/vm"
)

type Encoder struct {
	d *dumper.Dumper
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		d: dumper.NewDumper(lexer.NewUnlexer(w)),
	}
}

func (e *Encoder) Encode(v interface{}) error {
	val, err := types.ToValue(v)
	if err != nil {
		return err
	}
	return e.d.Dump(val)
}

type Decoder struct {
	l       *lexer.Lexer
	bufsize int
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		l: lexer.NewLexer(r),
	}
}

func (d *Decoder) SetBufSize(size int) {
	d.bufsize = size
}

func (d *Decoder) Decode(v interface{}) error {
	m := vm.NewVM(vm.WithStackSize(d.bufsize))
	for {
		tok, err := d.l.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		err = m.Feed(tok.Op)
		if err != nil {
			return err
		}
	}
	top, err := m.Top()
	if err != nil {
		return err
	}
	return top.Bind(v)
}
