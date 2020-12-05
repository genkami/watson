package vm

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeepCopyWithInt(t *testing.T) {
	orig := NewIntValue(123)
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Int = 456
	if orig.Int == clone.Int {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyString(t *testing.T) {
	orig := NewStringValue([]byte("hello"))
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.String[0] = 0x61 // 'a'
	if bytes.Equal(orig.String, clone.String) {
		t.Errorf("clone shares the same reference with its origin")
	}

	clone.String = []byte("world")
	if bytes.Equal(orig.String, clone.String) {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyWithObject(t *testing.T) {
	orig := NewObjectValue(map[string]*Value{
		"hello": NewStringValue([]byte("world")),
	})
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Object["hello"].String[0] = 0x61 // 'a'
	if diff := cmp.Diff(orig, clone); diff == "" {
		t.Errorf("clone shares the same reference with its origin")
	}

	clone.Object["hoge"] = NewStringValue([]byte("fuga"))
	if diff := cmp.Diff(orig, clone); diff == "" {
		t.Errorf("clone shares the same reference with its origin")
	}

	clone.Object = map[string]*Value{}
	if diff := cmp.Diff(orig, clone); diff == "" {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyWithBool(t *testing.T) {
	orig := NewBoolValue(true)
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Bool = false
	if orig.Bool == clone.Bool {
		t.Errorf("DeepCopy returned receiver itself")
	}
}

func TestDeepCopyWithNil(t *testing.T) {
	orig := NewNilValue()
	clone := orig.DeepCopy()
	if diff := cmp.Diff(orig, clone); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}

	clone.Kind = KInt
	if orig.Kind == clone.Kind {
		t.Errorf("DeepCopy returned receiver itself")
	}
}
