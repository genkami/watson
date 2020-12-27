package util

import (
	"flag"
	"fmt"
)

type Type int

const (
	Yaml Type = iota
	Json
)

const (
	typeNameYaml = "yaml"
	typeNameJson = "json"
)

func (t *Type) String() string {
	switch *t {
	case Yaml:
		return typeNameYaml
	case Json:
		return typeNameJson
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
	default:
		return fmt.Errorf("unknown type: %s", s)
	}
	return nil
}

var assertTypeIsValue = Type(0)
var _ flag.Value = &assertTypeIsValue
