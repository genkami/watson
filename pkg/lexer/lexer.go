package lexer

import (
	"io"

	"github.com/genkami/watson/pkg/repr"
	"github.com/genkami/watson/pkg/vm"
)

// Lexer is responsible for converting a Watson Representation into a sequence of vm.Ops.
type Lexer struct {
	r   io.Reader
	buf [1]byte
}

// Creates a new Lexer that reads Watson Representation from r.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{r: r}
}

// Returns the next Op.
// This returns io.EOF if it hits on the end of the input.
func (l *Lexer) Next() (vm.Op, error) {
	for {
		_, err := l.r.Read(l.buf[:])
		if err != nil {
			// Note that it returns io.EOF if the underlying Reader returns io.EOF.
			return 0, err
		}
		if op, ok := repr.ReadOp(l.buf[0]); ok {
			return op, nil
		}
	}
}
