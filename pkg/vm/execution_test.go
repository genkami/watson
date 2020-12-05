package vm

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFeedInewPushesZero(t *testing.T) {
	want := 2
	got := 1

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
