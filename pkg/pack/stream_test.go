//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStream — 40 documents (list and version) and a named record streamed with a crash and resume give a set that verifies, serves every path and h/ copy with the same bytes, and equals an uninterrupted run byte for byte
package pack

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStream(t *testing.T) {
	ctx := context.Background()
	finish := func(s *Stream) string {
		r, _ := s.Encode([]byte("q1 0x5100000000000001\n"))
		if err := s.PutNamed("0x00000001/codecs/0x00000420/0x00000001.geul", r); err != nil {
			t.Fatal(err)
		}
		res, err := s.Close()
		if err != nil || res.Records != 81 || res.Dict == 0 {
			t.Fatalf("%+v %v", res, err)
		}
		b, _ := os.ReadFile(filepath.Join(s.out, latestFile))
		return string(b)
	}
	// Uninterrupted.
	a := filepath.Join(t.TempDir(), "a")
	s, err := NewStream(a, streamTestOptions())
	if err != nil {
		t.Fatal(err)
	}
	streamTestPut(t, s, 1, 40)
	latestA := finish(s)
	// Crash after a checkpoint at document 15, with records put after it.
	b := filepath.Join(t.TempDir(), "b")
	s, _ = NewStream(b, streamTestOptions())
	streamTestPut(t, s, 1, 15)
	if err := s.Checkpoint([]byte("doc 15")); err != nil {
		t.Fatal(err)
	}
	streamTestPut(t, s, 16, 22)
	s.Checkpoint([]byte("doc 22")) // a later checkpoint survives too
	streamTestPut(t, s, 23, 26)
	if _, err := NewStream(b, streamTestOptions()); err == nil {
		t.Fatal("unfinished stream: NewStream must refuse")
	}
	s, err = ResumeStream(b, streamTestOptions())
	if err != nil || s.Count() != 44 || string(s.Meta()) != "doc 22" {
		t.Fatalf("resume: %d %v", s.Count(), err)
	}
	if err := s.Put([]core.Word{0, 1}, 0x5100000000000001, 0, StreamRecord{}); err == nil {
		t.Fatal("key before the last: want error")
	}
	streamTestPut(t, s, 23, 40)
	if latestB := finish(s); latestB != latestA {
		t.Fatalf("resumed run differs: %q vs %q", latestB, latestA)
	}
	if _, err := NewStream(b, streamTestOptions()); err == nil {
		t.Fatal("existing set: NewStream must refuse")
	}
	r := opened(t, b)
	if err := r.Verify(ctx); err != nil {
		t.Fatal(err)
	}
	var want []string
	for i := uint64(2); i < 82; i++ {
		doc := streamTestDoc(i/2, uint32(i%2))
		ws, _ := core.ReadWords(doc)
		h, _ := core.ContentHash(ws)
		want = append(want, pPath("0x00000001", 0x5100000000000000|i/2, uint32(i%2)), hPath("0x00000001", string(h[:])))
	}
	for i, path := range want {
		got, err := r.Fetch(ctx, path)
		if err != nil || !bytes.Equal(got, streamTestDoc(uint64(i/4+1), uint32(i/2%2))) {
			t.Fatalf("%s: %v", path, err)
		}
	}
	want = append(want, "0x00000001/codecs/0x00000420/0x00000001.geul")
	paths, err := r.Paths(ctx)
	sort.Strings(want)
	sort.Strings(paths)
	if err != nil || len(paths) != len(want) || paths[0] != want[0] || paths[len(paths)-1] != want[len(want)-1] {
		t.Fatalf("paths %d want %d %v", len(paths), len(want), err)
	}
	if _, err := r.Fetch(ctx, pPath("0x00000001", 0x5100000000000000|41, 1)); !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}
	if objs := len(r.m.objects); objs < 4 {
		t.Fatalf("expected several data packs and index objects, got %d", objs)
	}
}
