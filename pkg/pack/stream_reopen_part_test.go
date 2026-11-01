//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamReopenPart — none open is nil; the open object's part comes back cut to the checkpoint with its hash; one that cannot be found is an error
package pack

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamReopenPart(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	s.Checkpoint(nil)
	want := hashState(s.dsum)
	streamTestPut(t, s, 7, 7) // bytes after the checkpoint
	if f, h, err := s.reopenPart(streamDataPart, kindData, streamOpen{Num: -1}); f != nil || h != nil || err != nil {
		t.Fatal("none open")
	}
	b, _ := os.ReadFile(filepath.Join(s.out, streamDir, streamStateFile))
	st := streamTestState(t, b)
	r := &Stream{out: s.out, st: st}
	f, h, err := r.reopenPart(streamDataPart, kindData, st.Data)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if fi, _ := f.Stat(); uint64(fi.Size()) != st.Data.Off || !bytes.Equal(hashState(h), want) {
		t.Fatalf("size %d", fi.Size())
	}
	os.Remove(filepath.Join(s.out, streamDir, streamIndexPart))
	if _, _, err := r.reopenPart(streamIndexPart, kindIndex, st.Index); err == nil {
		t.Fatal("lost part: want error")
	}
}
