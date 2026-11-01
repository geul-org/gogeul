//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamFindPart — the part file is used when it matches; a matching unknown object in o/ is copied back over a missing or wrong part; neither matching is ErrCorrupt; an unreadable o/ or candidate is an error
package pack

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamFindPart(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	s.Checkpoint(nil)
	r, err := ResumeStream(s.out, streamTestOptions())
	if err != nil {
		t.Fatal(err)
	}
	part := filepath.Join(s.out, streamDir, streamDataPart)
	good, _ := os.ReadFile(part)
	f, err := r.findPart(streamDataPart, kindData, r.st.Data)
	if err != nil || f.Name() != part {
		t.Fatalf("part: %v", err)
	}
	f.Close()
	os.WriteFile(filepath.Join(s.out, "o", "sealed.gpk"), append(good, "tail"...), 0o644)
	os.WriteFile(part, []byte("a later object"), 0o644)
	f, err = r.findPart(streamDataPart, kindData, r.st.Data)
	if err != nil {
		t.Fatalf("sealed: %v", err)
	}
	f.Close()
	if got, _ := os.ReadFile(part); !bytes.Equal(got[headerSize:], good[headerSize:]) {
		t.Fatal("restored bytes differ")
	}
	os.Remove(filepath.Join(s.out, "o", "sealed.gpk"))
	os.Remove(part)
	if _, err := r.findPart(streamDataPart, kindData, r.st.Data); !errors.Is(err, ErrCorrupt) {
		t.Fatalf("neither: %v", err)
	}
	os.WriteFile(filepath.Join(s.out, "o", "locked.gpk"), nil, 0)
	if _, err := r.findPart(streamDataPart, kindData, r.st.Data); err == nil || errors.Is(err, ErrCorrupt) {
		t.Fatalf("unreadable candidate: %v", err)
	}
	os.RemoveAll(filepath.Join(s.out, "o"))
	if _, err := r.findPart(streamDataPart, kindData, r.st.Data); err == nil {
		t.Fatal("no o/: want error")
	}
	for _, rr := range []*Stream{r} {
		rr.data.Close()
		rr.index.Close()
	}
}
