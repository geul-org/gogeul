//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestLegacyIndexMatches — the P leaves a hash-less checkpoint wrote match back to back to its size and count; a wrong count, size, offset or leaf does not
package pack

import (
	"os"
	"testing"
)

func TestLegacyIndexMatches(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	f, _ := os.Open(s.index.Name())
	defer f.Close()
	open := s.st.Index
	if !legacyIndexMatches(f, open, s.st.PHeads) {
		t.Fatal("written leaves: want match")
	}
	for name, o := range map[string]streamOpen{"count": {Num: open.Num, Off: open.Off, Count: open.Count + 1}, "size": {Num: open.Num, Off: open.Off + 1, Count: open.Count}} {
		if legacyIndexMatches(f, o, s.st.PHeads) {
			t.Errorf("%s: want no match", name)
		}
	}
	heads := append([]streamKV{{Loc: Loc{Obj: 99}}}, s.st.PHeads...)
	heads[2].Loc.Off++
	if legacyIndexMatches(f, open, heads) {
		t.Fatal("shifted leaf: want no match")
	}
	heads = append([]streamKV(nil), s.st.PHeads...)
	heads[1].Key = "other"
	if legacyIndexMatches(f, open, heads) {
		t.Fatal("other leaf: want no match")
	}
}
