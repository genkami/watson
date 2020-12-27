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

var (
	outType util.Type
	mode    util.Mode
)

func buildFlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("watson encode", flag.ExitOnError)
	fs.Var(&outType, "t", "input type")
	fs.Var(&mode, "initial-mode", "initial mode of the lexer")
	return fs
}

func Main(args []string) {
	var err error
	fs := buildFlagSet()
	err = fs.Parse(args)
	if errors.Is(err, flag.ErrHelp) {
		os.Exit(0)
	} else if err != nil {
		fs.PrintDefaults()
		os.Exit(1)
	}

	lexer := lexer.NewLexer(os.Stdin, lexer.WithFileName("<stdin>"), lexer.WithInitialLexerMode(lexer.Mode(mode)))
	v, err := parseWatson(lexer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %s\n", err)
		os.Exit(1)
	}
	err = decode(os.Stdout, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't write Watson: %s\n", err.Error())
		os.Exit(1)
	}
}

type parseError struct {
	tok *lexer.Token
	err error
}

func (p *parseError) Error() string {
	return fmt.Sprintf("error %+v\n at %#v line %d, column %d\n",
		p.err, p.tok.FileName, p.tok.Line+1, p.tok.Column+1)
}

func parseWatson(lex *lexer.Lexer) (*types.Value, error) {
	m := vm.NewVM()
	for {
		tok, err := lex.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, &parseError{tok: tok, err: err}
		}
		err = m.Feed(tok.Op)
		if err != nil {
			return nil, &parseError{tok: tok, err: err}
		}
	}
	return m.Top()
}

func decode(w io.Writer, v *types.Value) error {
	switch outType {
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
