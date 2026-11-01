//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHEnd — no key no layer; a lone leaf is the root; otherwise the last leaf is written and the tree built over the head file; a failing leaf write, flush or tree is an error
package pack

import (
	"bufio"
	"os"
	"testing"
)

func TestStreamHEnd(t *testing.T) {
	s, p := streamTestHPass(t, nil)
	if ls, err := s.hEnd(p); ls != nil || err != nil {
		t.Fatal("no key")
	}
	var k [32]byte
	p.leaf, p.n = []streamKV{{Key: string(k[:])}}, 1
	if ls, err := s.hEnd(p); err != nil || ls[0].count != 1 || s.st.Index.Num != -1 {
		t.Fatal("lone leaf is the root")
	}
	full := func() (*Stream, *hPass) {
		s, p := streamTestHPass(t, nil)
		for i := 0; i < 5; i++ {
			k[0] = byte(i)
			p.leaf = []streamKV{{Key: string(k[:])}}
			s.hWriteLeaf(p)
		}
		k[0] = 9
		p.leaf, p.n = []streamKV{{Key: string(k[:])}}, 6
		return s, p
	}
	s, p = full()
	if ls, err := s.hEnd(p); err != nil || ls[0].count != 6 || p.count != 6 {
		t.Fatalf("tree: %v", err)
	}
	s, p = full()
	s.o.Build.Codec = failCodec{}
	if _, err := s.hEnd(p); err == nil {
		t.Fatal("failing leaf: want error")
	}
	s, p = full()
	s.o.Build.Codec = failCodecAfter(1)
	if _, err := s.hEnd(p); err == nil {
		t.Fatal("failing tree: want error")
	}
	s, p = full()
	dev, _ := os.OpenFile("/dev/full", os.O_WRONLY, 0)
	defer dev.Close()
	p.w = bufio.NewWriter(dev)
	if _, err := s.hEnd(p); err == nil {
		t.Fatal("failing flush: want error")
	}
}
