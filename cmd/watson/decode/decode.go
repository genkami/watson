package decode

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/genkami/watson/cmd/watson/util"
	"github.com/genkami/watson/pkg/converter/cbor"
	"github.com/genkami/watson/pkg/converter/json"
	"github.com/genkami/watson/pkg/converter/msgpack"
	"github.com/genkami/watson/pkg/converter/yaml"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/types"
	"github.com/genkami/watson/pkg/vm"
)

type Runner struct {
	outType util.Type
	mode    util.Mode
	files   []string
	m       *vm.VM
}

func NewRunner() *Runner {
	return &Runner{m: vm.NewVM()}
}

func (r *Runner) parseArgs(args []string) {
	fs := flag.NewFlagSet("watson decode", flag.ExitOnError)
	fs.Var(&r.outType, "t", "input type")
	fs.Var(&r.mode, "initial-mode", "initial mode of the lexer")
	err := fs.Parse(args)
	if errors.Is(err, flag.ErrHelp) {
		os.Exit(0)
	} else if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}
	r.files = fs.Args()
}

func (r *Runner) Run(args []string) {
	var err error
	r.parseArgs(args)

	lexer := r.buildLexer(os.Stdin, "<stdin>")
	v, err := r.parseWatson(lexer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %s\n", err)
		os.Exit(1)
	}
	err = r.decode(os.Stdout, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't write Watson: %s\n", err.Error())
		os.Exit(1)
	}
}

func (rn *Runner) buildLexer(r io.Reader, name string) *lexer.Lexer {
	return lexer.NewLexer(
		r,
		lexer.WithFileName(name),
		lexer.WithInitialLexerMode(lexer.Mode(rn.mode)),
	)
}

type parseError struct {
	tok *lexer.Token
	err error
}

func (p *parseError) Error() string {
	return fmt.Sprintf("error %+v\n at %#v line %d, column %d\n",
		p.err, p.tok.FileName, p.tok.Line+1, p.tok.Column+1)
}

func (r *Runner) parseWatson(lex *lexer.Lexer) (*types.Value, error) {
	for {
		tok, err := lex.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, &parseError{tok: tok, err: err}
		}
		err = r.m.Feed(tok.Op)
		if err != nil {
			return nil, &parseError{tok: tok, err: err}
		}
	}
	r.mode = util.Mode(lex.Mode())
	return r.m.Top()
}

func (r *Runner) decode(w io.Writer, v *types.Value) error {
	switch r.outType {
	case util.Yaml:
		return yaml.Decode(w, v)
	case util.Json:
		return json.Decode(w, v)
	case util.Msgpack:
		return msgpack.Decode(w, v)
	case util.Cbor:
		return cbor.Decode(w, v)
	default:
		panic("unknown output type")
	}
}
