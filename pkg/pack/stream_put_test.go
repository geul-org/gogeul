//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamPut — a closing stream acknowledges earlier keys and refuses later ones; records in key order fill leaves and spill H runs; an unknown Issuer, a key not after the last, a failed leaf write, place or spill are errors; a record without hash adds no H key
package pack

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStreamPut(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	if s.Count() != 12 || len(s.st.PHeads) != 2 || len(s.st.PLeaf) != 2 || len(s.st.Runs) != 1 || len(s.hbuf) != 5 || s.st.Res.Records != 12 {
		t.Fatalf("count %d heads %d leaf %d runs %d buf %d", s.Count(), len(s.st.PHeads), len(s.st.PLeaf), len(s.st.Runs), len(s.hbuf))
	}
	if err := s.Put([]core.Word{9}, 0x5100000000000007, 0, StreamRecord{}); err == nil {
		t.Fatal("unknown Issuer: want error")
	}
	if err := s.Put([]core.Word{1, 0}, 0x5100000000000006, 1, StreamRecord{}); err == nil {
		t.Fatal("same key: want error")
	}
	if err := s.Put([]core.Word{1, 0}, 0x5100000000000007, 0, StreamRecord{Bytes: []byte("x")}); err != nil || len(s.hbuf) != 5 {
		t.Fatalf("no hash: %v", err)
	}
	s.data.Close()
	if err := s.Put([]core.Word{1, 0}, 0x5100000000000007, 1, StreamRecord{Bytes: []byte("x")}); err == nil {
		t.Fatal("place failure: want error")
	}
	full := streamTestNew(t, nil)
	streamTestPut(t, full, 1, 2) // leaf of 4, one short of full
	r, _ := full.Encode(streamTestDoc(3, 0))
	full.Put([]core.Word{1, 0}, 0x5100000000000003, 0, r)
	os.RemoveAll(filepath.Join(full.out, streamDir))
	if err := full.Put([]core.Word{1, 0}, 0x5100000000000003, 1, r); err == nil {
		t.Fatal("leaf write failure: want error")
	}
	cl := streamTestNew(t, nil)
	streamTestPut(t, cl, 1, 2)
	cl.st.Close.Phase = closeStarted
	if err := cl.Put([]core.Word{1, 0}, 0x5100000000000001, 1, StreamRecord{}); err != nil || cl.Count() != 4 {
		t.Fatalf("closing, earlier key: %v", err)
	}
	if err := cl.Put([]core.Word{1, 0}, 0x5100000000000003, 0, StreamRecord{}); err == nil {
		t.Fatal("closing, later key: want error")
	}
	sp := streamTestNew(t, nil)
	streamTestPut(t, sp, 1, 3)
	os.RemoveAll(filepath.Join(sp.out, streamDir))
	r, _ = sp.Encode(streamTestDoc(4, 0))
	sp.data.Close()
	sp.st.Data.Num = -1
	sp.data = nil
	if err := sp.Put([]core.Word{1, 0}, 0x5100000000000004, 0, r); err == nil {
		t.Fatal("spill failure: want error")
	}
}
