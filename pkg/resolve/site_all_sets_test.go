//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteAllSets — a folder shows its Issuers; a server names them in issuers.geul, plus the sets opened so far; unreadable discovery is an error
package resolve

import (
	"context"
	"errors"
	"net/http"
	"path/filepath"
	"slices"
	"testing"
)

func TestSiteAllSets(t *testing.T) {
	ctx := context.Background()
	root := packedRoot(t, formsTree(t), false, nil)
	h := "h/0x00.geul"
	if got, err := localSite(root).allSets(ctx, h); err != nil || !slices.Equal(got, []string{"", "0x0000", "0x0001"}) {
		t.Fatal(got, err)
	}
	if _, err := localSite(filepath.Join(root, "0x0000", "latest.txt", "x")).allSets(ctx, h); err != nil {
		t.Fatal("no folder has no Issuers:", err)
	}
	notDir := filepath.Join(root, ".well-known", "geul", "0x0000", "latest.txt")
	file := &site{dir: notDir}
	if _, err := file.allSets(ctx, h); err == nil {
		t.Fatal("a file for a folder")
	}
	s, _ := served(t, root)
	if got, err := s.allSets(ctx, h); err != nil || !slices.Equal(got, []string{"", "0x0000", "0x0001"}) {
		t.Fatal(got, err)
	}
	if _, err := s.set(ctx, "0x0001"); err != nil {
		t.Fatal(err)
	}
	if got, err := s.allSets(ctx, issuersFile); err != nil || !slices.Equal(got, []string{"", "0x0000", "0x0001"}) {
		t.Fatal("0x0001 is open by now:", got, err)
	}
	fresh, _ := served(t, root)
	if got, err := fresh.allSets(ctx, issuersFile); err != nil || !slices.Equal(got, []string{"", "0x0000"}) {
		t.Fatal("issuers.geul itself is not looked up in issuers.geul:", got, err)
	}
	unnamed, _ := served(t, t.TempDir())
	if got, err := unnamed.allSets(ctx, h); err != nil || !slices.Equal(got, []string{"", "0x0000"}) {
		t.Fatal(got, err)
	}
	down, m := served(t, root)
	m.fail = func(w http.ResponseWriter, _ *http.Request) bool {
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}
	if _, err := down.allSets(ctx, h); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
