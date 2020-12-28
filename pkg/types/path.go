package types

import (
	"fmt"
)

type path interface {
	string() string
}

type rootPath struct{}

func newRootPath() path {
	return &rootPath{}
}

func (p *rootPath) string() string {
	return "<root>"
}

type fieldPath struct {
	parent path
	field  string
}

func newFieldPath(parent path, field string) path {
	return &fieldPath{
		parent: parent,
		field:  field,
	}
}

func (p *fieldPath) string() string {
	return fmt.Sprintf("%s.%s", p.parent.string(), p.field)
}

type indexPath struct {
	parent path
	idx    int
}

func newIndexPath(parent path, idx int) path {
	return &indexPath{
		parent: parent,
		idx:    idx,
	}
}
func (p *indexPath) string() string {
	return fmt.Sprintf("%s[%d]", p.parent.string(), p.idx)
}
