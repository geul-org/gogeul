//ff:func feature=cli type=command control=sequence
//ff:what TestNewDeprecateCmd — a version is marked deprecated; no version and no data folder fail
package cli

import (
	"strings"
	"testing"
)

func TestNewDeprecateCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	h.must("deprecate", addr+"@1")
	if out := h.must("show", "versions", addr); !strings.Contains(out, "deprecated") {
		t.Fatal(out)
	}
	if _, err := h.run("", "deprecate", addr); err == nil {
		t.Fatal("no version: want error")
	}
}
