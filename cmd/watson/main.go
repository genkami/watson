package main

import (
	"fmt"
	"os"

	"github.com/genkami/watson/cmd/watson/decode"
)

var allCmds = map[string]func([]string){
	"decode": decode.Main,
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd, ok := allCmds[os.Args[1]]
	if !ok {
		usage()
		os.Exit(1)
	}
	cmd(os.Args[2:])
}

func usage() {
	out := os.Stderr
	fmt.Fprintf(out, "usage: %s [", os.Args[0])
	first := true
	for name := range allCmds {
		if first {
			first = false
		} else {
			fmt.Fprintf(out, "|")
		}
		fmt.Fprintf(out, "%s", name)
	}
	fmt.Fprintf(out, "]\n")
}
