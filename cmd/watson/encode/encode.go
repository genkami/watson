package encode

import (
	"os"

	"github.com/genkami/watson/pkg/converter/yaml"
	"github.com/genkami/watson/pkg/dumper"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/prettifier"
)

func Main(args []string) {
	var err error
	val, err := yaml.Encode(os.Stdin)
	if err != nil {
		panic(err)
	}
	unl := prettifier.NewPrettifier(lexer.NewUnlexer(os.Stdout))
	d := dumper.NewDumper(unl)
	err = d.Dump(val)
	if err != nil {
		panic(err)
	}
}
