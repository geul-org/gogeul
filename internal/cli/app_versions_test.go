//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppVersions — states in text, JSON; a bad address, no data folder and an unpublished document fail
package cli

import (
	"strings"
	"testing"
)

func TestAppVersions(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	h.must("deprecate", addr+"@1")
	if out := h.must("show", "versions", addr); !strings.Contains(out, "\tdeprecated\n") {
		t.Fatal(out)
	}
	if out := h.must("show", "versions", addr, "--format", "json"); !strings.Contains(out, `"next_local_id": 4`) {
		t.Fatal(out)
	}
	for _, a := range []string{"bad", "me/0x0000000000000009"} {
		if _, err := h.run("", "show", "versions", a); err == nil {
			t.Errorf("%s: want error", a)
		}
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "show", "versions", "0x00000000/0x0000000000000001"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
