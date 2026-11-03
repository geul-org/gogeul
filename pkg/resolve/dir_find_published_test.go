//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirFindPublished — the 공표 주소 forms of the published sets read through GEULPROXY=file://…,off and the same bytes over HTTP Range; each set skipped until its latest.txt exists
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

// TestDirFindPublished reads the stand-in for geul.org/.well-known/geul/: the
// geul.org set 0x00000000 (codec documents, §7.1 step 5) and the Wikidata set
// 0x00000001 (§7.1 step 12), each only when its own `latest.txt` is there
// (migration.md §6.2). Single addresses only; no set is listed or walked. It is
// skipped when this module is checked out without geul-dev draft/ around it.
func TestDirFindPublished(t *testing.T) {
	root, _ := filepath.Abs("../../../geul-org/artifacts/geul")
	wk := filepath.Join(root, ".well-known", "geul")
	_, geulErr := os.Stat(filepath.Join(wk, "0x00000000", "latest.txt"))
	_, wikiErr := os.Stat(filepath.Join(wk, "0x00000001", "latest.txt"))
	if geulErr != nil && wikiErr != nil {
		t.Skip("geul-org/artifacts/geul/.well-known/geul holds neither 0x00000000 nor 0x00000001 yet")
	}
	c, err := ConfigFromEnv(func(k string) string {
		return map[string]string{"GEULPROXY": "file://" + root + ",off"}[k]
	})
	if err != nil {
		t.Fatal(err)
	}
	local := c.Proxy[0].Source.(Dir)
	s, meter := served(t, root)
	remote := Dir{Store: s}
	var addrs, missing []string
	if geulErr == nil {
		codec := "0x00000000/codecs/0x00000980"
		v1, err := local.Find(codec + "/0x00000001.geul")
		if err != nil {
			t.Fatal(err)
		}
		ws, _ := core.ReadWords(v1)
		sum, _ := core.ContentHash(ws)
		addrs = append(addrs, "issuers/latest.geul", codec+"/latest.geul", codec+"/0x00000001.geul", codec+"/versions.geul", hashPath(geulIssuer, sum))
		missing = append(missing, "0x00000000/codecs/0x0000fff0/latest.geul", "0x00000000/hashes/0x"+strings.Repeat("0", 64)+".geul",
			"0x00000000/docs/latest.geul", "0x00000000/docs/versions.geul", "0x00000000/codecs/latest.geul", "0x00000000/codecs/versions.geul")
	}
	if wikiErr == nil {
		item := "0x00000001/docs/0x510000000000002a"
		addrs = append(addrs, item+"/latest.geul", item+"/0x00000001.geul", item+"/versions.geul", "0x0000000000000001/docs/0x510000000000002a/latest.geul")
		missing = append(missing, "0x00000001/docs/0x51000000ffffffff/latest.geul")
	}
	for _, a := range addrs {
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
	for _, a := range missing {
		if _, err := local.Find(a); !errors.Is(err, ErrNotFound) {
			t.Fatal(a, err)
		}
		if _, err := remote.Find(a); !errors.Is(err, ErrNotFound) {
			t.Fatal(a, err)
		}
	}
}
