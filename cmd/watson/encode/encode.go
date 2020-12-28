package encode

import (
	"errors"
	"flag"
	"io"
	"os"

	"github.com/genkami/watson/cmd/watson/util"
	"github.com/genkami/watson/pkg/converter/cbor"
	"github.com/genkami/watson/pkg/converter/json"
	"github.com/genkami/watson/pkg/converter/msgpack"
	"github.com/genkami/watson/pkg/converter/yaml"
	"github.com/genkami/watson/pkg/dumper"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/prettifier"
	"github.com/genkami/watson/pkg/types"
)

type Runner struct {
	inType util.Type
	mode   util.Mode
	files  []string
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) parseArgs(args []string) {
	fs := flag.NewFlagSet("watson encode", flag.ExitOnError)
	fs.Var(&r.inType, "t", "input type")
	fs.Var(&r.mode, "initial-mode", "initial mode of the unlexer")
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
	val, err := r.encode(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = r.dump(os.Stdout, val)
	if err != nil {
		panic(err)
	}
}

func (rn *Runner) encode(r io.Reader) (*types.Value, error) {
	switch rn.inType {
	case util.Yaml:
		return yaml.Encode(r)
	case util.Json:
		return json.Encode(r)
	case util.Msgpack:
		return msgpack.Encode(r)
	case util.Cbor:
		return cbor.Encode(r)
	default:
		panic("unknown input type")
	}
}

func (r *Runner) dump(w io.Writer, v *types.Value) error {
	unl := prettifier.NewPrettifier(lexer.NewUnlexer(w, lexer.WithInitialUnlexerMode(lexer.Mode(r.mode))))
	d := dumper.NewDumper(unl)
	return d.Dump(v)
}
