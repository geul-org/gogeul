//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHWriteLeaf — the leaf is written as an H leaf block and its head appended to the head file; a failing block write is an error
package pack

import (
	"testing"
)

func TestStreamHWriteLeaf(t *testing.T) {
	s, p := streamTestHPass(t, nil)
	var k [32]byte
	k[0] = 7
	p.leaf = []streamKV{{Key: string(k[:]), Loc: Loc{Off: 9}}}
	if err := s.hWriteLeaf(p); err != nil || p.count != 1 || p.since != 1 || len(p.leaf) != 0 || s.st.Index.Count != 1 {
		t.Fatal(err)
	}
	p.w.Flush()
	heads, err := readHeads(p.heads, 0, 1)
	if err != nil || heads[0].Key != string(k[:]) || heads[0].Loc.Obj != uint32(s.st.Index.Num) {
		t.Fatalf("%v %v", heads, err)
	}
	f, fp := streamTestHPass(t, failCodec{})
	fp.leaf = []streamKV{{Key: string(k[:])}}
	if err := f.hWriteLeaf(fp); err == nil {
		t.Fatal("failing block: want error")
	}
}
