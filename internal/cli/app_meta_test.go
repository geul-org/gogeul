//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppMeta — a replaced Meta keeps its place; an unknown name, no value, a missing draft and a value breaking the spec fail
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestAppMeta(t *testing.T) {
	h := newHarness(t)
	_, draft := h.doc()
	h.must("meta", draft, "SPEC-VERSION", "0x0002")
	if out := h.must("show", draft); !strings.HasPrefix(out, "SPEC-VERSION 0x0002\nDOC-ID") {
		t.Fatal(out)
	}
	cases := [][]string{{draft, "NOPE", "1"}, {draft, "MESSAGE"}, {filepath.Join(h.data, "none"), "MESSAGE", `"a"`}, {draft, "NEXT-LOCAL-ID", "0x00000000"}}
	for _, c := range cases {
		if _, err := h.run("", append([]string{"meta"}, c...)...); err == nil {
			t.Errorf("%v: want error", c)
		}
	}
}
