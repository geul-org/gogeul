//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestNewGetCmd — geul get writes the stored bytes of a file:// 공표 주소 (latest, a version, the version list, a padded Issuer); other forms and missing addresses fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestNewGetCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	base := "file://" + h.data + "/.well-known/geul/"
	k, _ := h.key(addr)
	folder := filepath.Join(h.data, ".well-known", "geul", core.FormatID(k.Issuer), "docs", core.FormatID(k.Doc))
	v1, _ := os.ReadFile(filepath.Join(folder, "0x00000001.geul"))
	v2, _ := os.ReadFile(filepath.Join(folder, "0x00000002.geul"))
	list, _ := os.ReadFile(filepath.Join(folder, "versions.geul"))
	rel := strings.TrimPrefix(folder, filepath.Join(h.data, ".well-known", "geul")+"/")
	for a, want := range map[string][]byte{rel + "/latest.geul": v2, rel + "/0x00000002.geul": v2, rel + "/0x00000001.geul": v1, rel + "/versions.geul": list} {
		if out := h.must("get", base+a); out != string(want) || len(want) == 0 {
			t.Fatalf("%s: %d bytes, want %d", a, len(out), len(want))
		}
	}
	for _, a := range []string{rel + "/0x00000009.geul", rel + "/latest", "issuers.geul", rel + ".geul", "me/docs/0x0000000000000001/latest.geul"} {
		if _, err := h.run("", "get", base+a); err == nil || !strings.Contains(err.Error(), "not found") {
			t.Fatalf("%s: %v", a, err)
		}
	}
	if _, err := getAddr(nil, "0x00000000/docs/latest.geul"); err == nil {
		t.Fatal("no base: want error")
	}
	if _, err := resolve.ParseAddr(base + rel + "/latest.geul"); err != nil {
		t.Fatal(err)
	}
}
