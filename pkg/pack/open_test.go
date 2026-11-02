//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestOpen — latest.txt and pinned opens; missing or bad latest, a latest.txt naming a missing manifest (ErrCorrupt, not NotFound), missing, renamed, unparsable or mismatched manifests, rollback and bad roots fail
package pack

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpen(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	if r.Generation() != 1 {
		t.Fatal(r.Generation())
	}
	s := FileStore{Root: out}
	if _, err := Open(ctx, s, Options{Manifest: r.ManifestName()}); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(ctx, s, Options{MinGeneration: 2}); !errors.Is(err, ErrRollback) {
		t.Fatal(err)
	}
	if _, err := Open(ctx, FileStore{Root: t.TempDir()}, Options{}); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := Open(ctx, s, Options{Manifest: strings.Repeat("0", 64)}); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	m := memStore{latestFile: []byte("x")}
	if _, err := Open(ctx, m, Options{}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := Open(ctx, memStore{latestFile: []byte("1 " + strings.Repeat("0", 64) + "\n")}, Options{}); !errors.Is(err, ErrCorrupt) || errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	mb, _ := os.ReadFile(filepath.Join(out, "m", r.ManifestName()+".gpm"))
	if _, err := Open(ctx, memStore{"m/" + strings.Repeat("0", 64) + ".gpm": mb}, Options{Manifest: strings.Repeat("0", 64)}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	junk := sha256.Sum256([]byte("junk"))
	name2 := hex.EncodeToString(junk[:])
	if _, err := Open(ctx, memStore{"m/" + name2 + ".gpm": []byte("junk")}, Options{Manifest: name2}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	bad, _ := memGen(&manifest{gen: 1, objects: []object{}}, nil)
	for k := range bad {
		if strings.HasPrefix(k, "m/") {
			bad[latestFile] = []byte("2 " + strings.TrimSuffix(strings.TrimPrefix(k, "m/"), ".gpm") + "\n")
		}
	}
	if _, err := Open(ctx, bad, Options{}); !errors.Is(err, ErrCorrupt) {
		t.Fatal("generation", err)
	}
	pBad, _ := memGen(&manifest{gen: 1, pLayers: []layer{{enc: 0, root: []byte{9}}}}, nil)
	if _, err := Open(ctx, pBad, Options{}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	hBad, _ := memGen(&manifest{gen: 1, hLayers: []layer{{enc: 0, root: []byte{9}}}}, nil)
	if _, err := Open(ctx, hBad, Options{}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
