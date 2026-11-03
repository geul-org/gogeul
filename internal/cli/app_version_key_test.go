//ff:func feature=cli type=helper control=sequence
//ff:what TestAppVersionKey — a version is required; no data folder fails before resolving
package cli

import "testing"

func TestAppVersionKey(t *testing.T) {
	h := newHarness(t)
	if _, err := h.run("", "deprecate", "0x00000000/0x0000000000000001"); err == nil {
		t.Fatal("no version: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "deprecate", "0x00000000/0x0000000000000001@1"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
