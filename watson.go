package watson

import (
	"bytes"
	"io"

	"github.com/genkami/watson/pkg/dumper"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/types"
	"github.com/genkami/watson/pkg/vm"
)

func Marshal(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := NewEncoder(buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(buf []byte, v interface{}) error {
	dec := NewDecoder(bytes.NewReader(buf))
	return dec.Decode(v)
}

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
	l         *lexer.Lexer
	stackSize int
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		l: lexer.NewLexer(r),
	}
}

func (d *Decoder) SetStacksize(size int) {
	d.stackSize = size
}

func (d *Decoder) Decode(v interface{}) error {
	m := vm.NewVM(vm.WithStackSize(d.stackSize))
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
