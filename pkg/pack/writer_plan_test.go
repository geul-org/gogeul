//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterPlan — duplicates collapse, unchanged paths drop, changed ones stay; Get errors and sharing errors fail
package pack

import (
	"errors"
	"strings"
	"testing"
)

func TestWriterPlan(t *testing.T) {
	_, out := built(t, small())
	get := memGet(map[string][]byte{"issuers/0x00000001.geul": fixtureFile(t, "issuers/0x00000001.geul"), "issuers/0x00000002.geul": {1}})
	w := testWriter(t, out, Batch{Put: []string{"issuers/0x00000002.geul", "issuers/0x00000001.geul", "issuers/0x00000002.geul"}, Get: get})
	if err := w.plan(); err != nil || len(w.puts) != 1 || w.puts[0].path != "issuers/0x00000002.geul" || len(w.sums) != 2 {
		t.Fatal(err, w.puts)
	}
	w = testWriter(t, out, Batch{Put: []string{"y"}, Get: get})
	if err := w.plan(); err == nil {
		t.Fatal("get")
	}
	h := "0x00000000/hashes/0x" + strings.Repeat("0", 64) + ".geul"
	w = testWriter(t, out, Batch{Put: []string{h}, Get: memGet(map[string][]byte{h: {1, 2}})})
	if err := w.plan(); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
