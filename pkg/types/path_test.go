package types

import (
	"testing"
)

func TestRootPath(t *testing.T) {
	path := newRootPath()
	expected := "<root>"
	actual := path.string()
	if expected != actual {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestFieldPath(t *testing.T) {
	path := newFieldPath(newRootPath(), "TheField")
	expected := "<root>.TheField"
	actual := path.string()
	if expected != actual {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}

func TestIndexPath(t *testing.T) {
	path := newIndexPath(newFieldPath(newRootPath(), "TheField"), 1)
	expected := "<root>.TheField[1]"
	actual := path.string()
	if expected != actual {
		t.Errorf("expected %#v but got %#v", expected, actual)
	}
}
