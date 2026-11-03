//ff:func feature=pack type=writer control=sequence
//ff:what TestListTree — the fixture's 18 files sorted; a missing archive fails
package pack

import (
	"sort"
	"testing"
)

func TestListTree(t *testing.T) {
	ps, err := listTree("testdata/archive")
	if err != nil || len(ps) != 18 || !sort.StringsAreSorted(ps) || ps[len(ps)-1] != "issuers/0x00000001.geul" {
		t.Fatal(ps, err)
	}
	if _, err := listTree(t.TempDir()); err == nil {
		t.Fatal("missing")
	}
}
