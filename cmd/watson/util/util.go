package util

import (
	"flag"
	"fmt"
	"io"
	"os"

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

type Opener interface {
	Name() string
	Open() (io.ReadWriteCloser, error)
}

type RWCOpener struct {
	name string
	rwc  io.ReadWriteCloser
}

func NewRWCOpener(name string, rwc io.ReadWriteCloser) *RWCOpener {
	return &RWCOpener{
		name: name,
		rwc:  rwc,
	}
}

func (rwc *RWCOpener) Name() string {
	return rwc.name
}

func (rwc *RWCOpener) Open() (io.ReadWriteCloser, error) {
	return rwc.rwc, nil
}

var _ Opener = &RWCOpener{}

type FileOpener struct {
	path string
	flag int
	perm os.FileMode
}

func NewFileOpener(path string, flag int, perm os.FileMode) *FileOpener {
	return &FileOpener{
		path: path,
		flag: flag,
		perm: perm,
	}
}

func (o *FileOpener) Name() string {
	return o.path
}

func (o *FileOpener) Open() (io.ReadWriteCloser, error) {
	return os.OpenFile(o.path, o.flag, o.perm)
}

var _ Opener = &FileOpener{}
