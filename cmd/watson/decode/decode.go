package decode

import (
	"fmt"
	"io"
	"os"

	"github.com/genkami/watson/pkg/converter/yaml"
	"github.com/genkami/watson/pkg/lexer"
	"github.com/genkami/watson/pkg/vm"
)

func Main(args []string) {
	m := vm.NewVM()
	lexer := lexer.NewLexer(os.Stdin, lexer.WithFileName("<stdin>"))
	for {
		tok, err := lexer.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			showErrAndExit(tok, err)
		}
		err = m.Feed(tok.Op)
		if err != nil {
			showErrAndExit(tok, err)
		}
	}
	v, err := m.Top()
	if err != nil {
		panic(err)
	}
	err = yaml.Decode(os.Stdout, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "output is empty")
		os.Exit(1)
	}
}

func showErrAndExit(tok *lexer.Token, err error) {
	fmt.Fprintf(os.Stderr, "error %+v\n at %#v line %d, column %d\n",
		err, tok.FileName, tok.Line+1, tok.Column+1)
	os.Exit(1)
}
