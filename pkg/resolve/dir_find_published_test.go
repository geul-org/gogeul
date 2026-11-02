//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirFindPublished — every kind of 공표 주소 of the published pack archive reads through GEULPROXY=file://…,off, and the same bytes come over HTTP Range, within a request budget
package resolve

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestDirFindPublished reads the stand-in for geul.org/.well-known/geul/. It
// is skipped when this module is checked out without geul-dev draft/ around it.
func TestDirFindPublished(t *testing.T) {
	root, _ := filepath.Abs("../../../geul-org/artifacts/geul")
	if _, err := os.Stat(filepath.Join(root, ".well-known", "geul", "0x0000", "latest.txt")); err != nil {
		t.Skip("geul-org/artifacts/geul/.well-known/geul holds no pack sets")
	}
	c, err := ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	if err != nil {
		t.Fatal(err)
	}
	local := c.Proxy[0].Source.(Dir)
	spec := "0x0000/0x0000000000001018"
	item := "0x00010000/0x510000000000002a"
	v1, err := local.Find(spec + "/0x00000001.geul")
	if err != nil {
		t.Fatal(err)
	}
	ws, _ := core.ReadWords(v1)
	sum, _ := core.ContentHash(ws)
	s, meter := served(t, root)
	remote := Dir{Store: s}
	for _, a := range []string{"issuers.geul", "0x0000/documents.geul", spec + ".geul", spec + "/0x00000001.geul", spec + "/versions.geul", hashPath(sum), item + ".geul", item + "/versions.geul"} {
		want, err := local.Find(a)
		if err != nil {
			t.Fatal(a, err)
		}
		got, err := remote.Find(a)
		if err != nil || !bytes.Equal(got, want) {
			t.Fatal(a, err)
		}
		seen, sent := meter.take()
		t.Logf("%-60s %7d bytes  https: %2d requests, %8d bytes", a, len(want), len(seen), sent)
	}
	missing := "0x0000/0x000000000000ffff.geul"
	if _, err := local.Find(missing); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := remote.Find(missing); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if h := strings.Repeat("0", 64); h != "" {
		if _, err := remote.Find("h/0x" + h + ".geul"); !errors.Is(err, ErrNotFound) {
			t.Fatal(err)
		}
	}
	seen, sent := meter.take()
	t.Logf("%-60s %7d bytes  https: %2d requests, %8d bytes", "2 misses", 0, len(seen), sent)
}
