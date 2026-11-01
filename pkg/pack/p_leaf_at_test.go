//ff:func feature=pack type=decoder control=sequence
//ff:what TestPLeafAt — a written P leaf matches its head; another first key, a failing read, a frame that is no zstd and an H leaf do not
package pack

import (
	"os"
	"testing"
)

func TestPLeafAt(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	h := s.st.PHeads[0]
	f, _ := os.Open(s.index.Name())
	defer f.Close()
	if !pLeafAt(f, h) {
		t.Fatal("written leaf: want match")
	}
	if pLeafAt(f, streamKV{Key: "other", Loc: h.Loc}) {
		t.Fatal("other key: want no match")
	}
	if pLeafAt(&badFile{readAt: true}, h) {
		t.Fatal("failing read: want no match")
	}
	if pLeafAt(f, streamKV{Key: h.Key, Loc: Loc{Off: 1, Len: 10}}) {
		t.Fatal("not a frame: want no match")
	}
	var e hEntry
	loc, _ := s.addBlock(blockOf(blkHLeaf, []streamKV{{Key: string(e.key[:])}}))
	if pLeafAt(f, streamKV{Key: string(e.key[:]), Loc: loc}) {
		t.Fatal("H leaf: want no match")
	}
}
