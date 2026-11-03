//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestRemoteSite — request budgets: a set is found once per Issuer, documents take only ranges of objects, never a whole object; a tree of files costs two probes then one GET; a server ignoring Range is refused
package resolve

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestRemoteSite(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	list, version := listPath(sampleKeyOf(7)), versionPath(sampleKeyOf(7), 2)
	s, m := served(t, packedRoot(t, tree, false, &pack.Zstd{Level: 1}))
	if _, err := s.Fetch(ctx, list); err != nil {
		t.Fatal(err)
	}
	seen, _ := m.take()
	if len(seen) < 3 || seen[0] != "0x00000000/latest.txt " || !strings.HasPrefix(seen[1], "0x00000000/m/") {
		t.Fatal(seen)
	}
	for _, r := range seen[2:] {
		if !strings.Contains(r, " bytes=") {
			t.Fatal("whole object:", r)
		}
	}
	if _, err := s.Fetch(ctx, version); err != nil {
		t.Fatal(err)
	}
	if seen, _ = m.take(); len(seen) != 1 || !strings.Contains(seen[0], " bytes=") {
		t.Fatal("a second document in the set is one range:", seen)
	}
	if _, err := s.Fetch(ctx, "0x00000000/docs/0xffff000000000009/versions.geul"); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	seen, _ = m.take()
	if len(seen) != 2 || seen[0] != "latest.txt " || seen[1] != "0x00000000/docs/0xffff000000000009/versions.geul " {
		t.Fatal("a miss asks for the top set once and the loose file:", seen)
	}
	s, m = served(t, tree.Root)
	for _, p := range []string{list, version, list} {
		if _, err := s.Fetch(ctx, p); err != nil {
			t.Fatal(err)
		}
	}
	seen, _ = m.take()
	if len(seen) != 5 || seen[0] != "0x00000000/latest.txt " || seen[1] != "latest.txt " {
		t.Fatal(seen)
	}
	s, m = served(t, packedRoot(t, tree, false, nil))
	m.fail = func(w http.ResponseWriter, r *http.Request) bool {
		if r.Header.Get("Range") == "" {
			return false
		}
		r.Header.Del("Range")
		m.next.ServeHTTP(w, r)
		return true
	}
	if _, err := s.Fetch(ctx, version); !errors.Is(err, pack.ErrRange) {
		t.Fatal(err)
	}
}
