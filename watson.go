// Package watson provides a convenient way of encoding and decoding Watson values.
package watson

import (
	"bytes"
	"io"

	"github.com/genkami/watson/pkg/dumper"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/types"
	"github.com/genkami/watson/pkg/vm"
)

// Marshal converts an arbitrary value `v` into Watson by applying the following rules:
//   * If v is any of int, int8, int16, int32, or int64, then v is converted to Int.
//   * If v is any of uint, uint8, uint16, uint32, or uint64, then v is converted to Uint.
//   * If v is float32 or float64, then v is converted to Float.
//   * If v is bool, then v is converted to Bool.
//   * If v is string, then v is converted to String.
//   * If v is a struct that implements `types.Marshaler`, then v is converted to Value by calling `v.MarshalWatson()`.
//   * If v is a struct that does not implement `types.Marshaler`, then v is converted to Object with its keys correspond to the fields of v.
//   * If v is a slice or an array, then v is converted to Array with its elements converted by these rules.
//   * If v is a map, then v is converted to Object with its elements converted by these rules.
//   * If v is a pointer, then v is converted to `Value` by converting `*v` with these rules.
//
// Note that you can configure struct fields by adding "watson" tag to fields.
// Tag must be like `watson:"name,flag1,flag2,...,flagN"`.
// If Marshal finds a field that has such tag, it uses `name` as a key of output instead of using the name of the field, or omits such field if `name` equals to "-".
//
// Currntly these flags are available:
//   omitempty      If the field is zero value, it will be omitted from the output.
//   inline         Inline the field. Currently the field must be a struct.
func Marshal(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := NewEncoder(buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshal converts Watson into an arbitrary object.
//
// You can customize its behavior by implementing `types.Unmarshaler`.
//
// See Marshal for details.
func Unmarshal(buf []byte, v interface{}) error {
	dec := NewDecoder(bytes.NewReader(buf))
	return dec.Decode(v)
}

// Encoder writes Watson values to a given io.Writer.
type Encoder struct {
	d *dumper.Dumper
}

// NewEncoder creates a new Encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		d: dumper.NewDumper(lexer.NewUnlexer(w)),
	}
}

// Encode writes the Watson encoding of v to the underlying io.Writer.
func (e *Encoder) Encode(v interface{}) error {
	val, err := types.ToValue(v)
	if err != nil {
		return err
	}
	return e.d.Dump(val)
}

// Decoder reads and decodes Watson values from a given io.Reader.
type Decoder struct {
	l         *lexer.Lexer
	stackSize int
}

// NewDecoder creates a new Decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		l: lexer.NewLexer(r),
	}
}

// SetStackSize sets the stack size of underlying Watson VM.
//
// See watson/pkg/vm for more details.
func (d *Decoder) SetStacksize(size int) {
	d.stackSize = size
}

// Decode reads a Watson value from the underlying io.Reader and converts it into v.
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
