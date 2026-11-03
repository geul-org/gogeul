//ff:func feature=cli type=helper control=sequence
//ff:what TestPackSources — GEULPACK items come before the GEULPROXY ones, which stay as written and open nothing until read; bad items fail
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestPackSources(t *testing.T) {
	h := newHarness(t)
	h.doc()
	plain := h.data
	h.packed()
	cfg, _ := resolve.ConfigFromEnv(func(k string) string {
		if k == "GEULPROXY" {
			return "file://" + h.data + ",file://" + plain + ",registry"
		}
		return ""
	})
	set := filepath.Join(h.data, ".well-known", "geul", "0x00000000")
	cfg, err := packSources(cfg, " "+set+" ,,")
	if err != nil || len(cfg.Proxy) != 4 || cfg.Proxy[0].Name != "pack+"+set || cfg.Proxy[1].Name != "file://"+h.data {
		t.Fatal(cfg.Proxy, err)
	}
	if _, err := packSources(resolve.Config{}, t.TempDir()); err == nil {
		t.Fatal("bad GEULPACK item")
	}
	os.WriteFile(filepath.Join(set, "latest.txt"), []byte("x"), 0o644)
	cfg, _ = resolve.ConfigFromEnv(func(k string) string {
		if k == "GEULPROXY" {
			return "file://" + h.data
		}
		return ""
	})
	if _, err := packSources(cfg, ""); err != nil {
		t.Fatal("a broken GEULPROXY folder fails when it is read, so `|` can pass it:", err)
	}
}
