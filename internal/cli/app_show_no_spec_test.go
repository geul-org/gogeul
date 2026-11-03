//ff:func feature=cli type=command control=sequence
//ff:what TestAppShowNoSpec — a codec without a published spec is shown generically, not refused
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppShowNoSpec(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	with := h.must("show", docsAddr(addr)+"/latest")
	os.RemoveAll(filepath.Join(h.data, ".well-known", "geul", "0x00000000", "codecs", "0x00000AA0"))
	out := h.must("show", docsAddr(addr)+"/latest")
	if strings.Contains(out, "normal") || !strings.Contains(out, "TRIPLE") || out == with {
		t.Fatal(out)
	}
}
