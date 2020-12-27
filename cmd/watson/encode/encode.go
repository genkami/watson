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

var (
	inType util.Type
	mode   util.Mode
)

func buildFlagSet() *flag.FlagSet {
	fs := flag.NewFlagSet("watson encode", flag.ExitOnError)
	fs.Var(&inType, "t", "input type")
	fs.Var(&mode, "initial-mode", "initial mode of the unlexer")
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

	val, err := encode(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = dump(os.Stdout, val)
	if err != nil {
		panic(err)
	}
}

func encode(r io.Reader) (*types.Value, error) {
	switch inType {
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

func dump(w io.Writer, v *types.Value) error {
	unl := prettifier.NewPrettifier(lexer.NewUnlexer(w, lexer.WithInitialUnlexerMode(lexer.Mode(mode))))
	d := dumper.NewDumper(unl)
	return d.Dump(v)
}
