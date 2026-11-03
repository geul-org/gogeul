//ff:func feature=pack type=helper control=sequence
//ff:what TestCanonSegs — every segment of a split path is rewritten in place by canonSeg
package pack

import (
	"slices"
	"testing"
)

func TestCanonSegs(t *testing.T) {
	seg := []string{"0x0000000a", "codecs", "0x00000abc", "0x0000000b.geul"}
	got := canonSegs(seg)
	want := []string{"0x0000000A", "codecs", "0x00000ABC", "0x0000000B.geul"}
	if !slices.Equal(got, want) || !slices.Equal(seg, want) {
		t.Fatal(got, seg)
	}
	if len(canonSegs(nil)) != 0 {
		t.Fatal("nil")
	}
}
