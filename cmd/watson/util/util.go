package util

import (
	"flag"
	"fmt"

	"github.com/genkami/watson/pkg/lexer"
)

type Mode lexer.Mode

const (
	modeNameA = "A"
	modeNameS = "S"
)

func (m *Mode) String() string {
	switch lexer.Mode(*m) {
	case lexer.A:
		return modeNameA
	case lexer.S:
		return modeNameS
	default:
		panic("unknown mode")
	}
}

func (m *Mode) Set(s string) error {
	switch s {
	case "":
		*m = Mode(lexer.A)
	case modeNameA:
		*m = Mode(lexer.A)
	case modeNameS:
		*m = Mode(lexer.S)
	default:
		return fmt.Errorf("unknown mode: %s", s)
	}
	return nil
}

var assertModeIsValue = Mode(0)
var _ flag.Value = &assertModeIsValue

type Type int

const (
	Yaml Type = iota
	Json
	Msgpack
	Cbor
)

const (
	typeNameYaml    = "yaml"
	typeNameJson    = "json"
	typeNameMsgpack = "msgpack"
	typeNameCbor    = "cbor"
)

func (t *Type) String() string {
	switch *t {
	case Yaml:
		return typeNameYaml
	case Json:
		return typeNameJson
	case Msgpack:
		return typeNameMsgpack
	case Cbor:
		return typeNameCbor
	default:
		panic("unknown type")
	}
}

func (t *Type) Set(s string) error {
	switch s {
	case "":
		*t = Yaml
	case typeNameYaml:
		*t = Yaml
	case typeNameJson:
		*t = Json
	case typeNameMsgpack:
		*t = Msgpack
	case typeNameCbor:
		*t = Cbor
	default:
		return fmt.Errorf("unknown type: %s", s)
	}
	return nil
}

var assertTypeIsValue = Type(0)
var _ flag.Value = &assertTypeIsValue
