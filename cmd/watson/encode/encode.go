package encode

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
	"github.com/genkami/watson/pkg/dumper"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/prettifier"
	"github.com/genkami/watson/pkg/types"
)

type Runner struct {
	inType util.Type
	mode   util.Mode
	opener util.Opener
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
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		fs.PrintDefaults()
		os.Exit(1)
	}
	files := fs.Args()
	if len(files) == 0 {
		r.opener = util.NewRWCOpener("<stdin>", os.Stdin)
	} else if len(files) == 1 {
		r.opener = util.NewFileOpener(files[0], os.O_RDONLY, 0)
	} else {
		fmt.Fprintf(os.Stderr, "too many arguments")
		fs.PrintDefaults()
		os.Exit(1)
	}
}

func (r *Runner) Run(args []string) {
	var err error
	r.parseArgs(args)
	file, err := r.opener.Open()
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open %s: %s\n", r.opener.Name(), err.Error())
		os.Exit(1)
	}
	defer file.Close()
	val, err := r.encode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error loading %s: %s\n", r.opener.Name(), err.Error())
		os.Exit(1)
	}
	err = r.dump(os.Stdout, val)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error writing output: %s\n", r.opener.Name(), err.Error())
		os.Exit(1)
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
