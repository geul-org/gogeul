//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestAppShowNoListing — show, a missing document or version and show versions read a few blocks of a set whose listing reads many; a full walk fails the budget
package cli

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestAppShowNoListing(t *testing.T) {
	const budget = 20
	ctx := context.Background()
	h := newHarness(t)
	addr, draft := h.doc()
	wk := filepath.Join(h.data, ".well-known", "geul", "0x0000")
	for i := range 300 {
		f := filepath.Join(wk, fmt.Sprintf("0x00001000%08x", i), "versions.geul")
		os.MkdirAll(filepath.Dir(f), 0o755)
		os.WriteFile(f, []byte{0, 0}, 0o644)
	}
	set := t.TempDir()
	small := pack.BuildOptions{PLeaf: 2, HLeaf: 2, HFanout: 2, HRoot: 2}
	if _, err := pack.Build(ctx, h.data, set, small); err != nil {
		t.Fatal(err)
	}
	m := &meterStore{Store: pack.FileStore{Root: set}}
	arch, err := pack.OpenRemote(ctx, m, pack.Options{})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := arch.Paths(); err != nil || m.reads <= 5*budget {
		t.Fatalf("the meter does not see a listing: %d reads, %v", m.reads, err)
	}
	missing := strings.TrimSuffix(addr, "1") + "9"
	for _, c := range []struct {
		args []string
		fail bool
	}{
		{[]string{"show", addr}, false},
		{[]string{"show", addr + "@1"}, false},
		{[]string{"show", draft}, false},
		{[]string{"show", "versions", addr}, false},
		{[]string{"show", missing}, true},
		{[]string{"show", addr + "@9"}, true},
		{[]string{"show", "versions", missing}, true},
	} {
		reads, err := h.metered(set, c.args...)
		if (err != nil) != c.fail {
			t.Fatal(c.args, err)
		}
		if reads > budget {
			t.Fatalf("%v: %d range reads, budget %d", c.args, reads, budget)
		}
		t.Log(c.args, reads)
	}
}
