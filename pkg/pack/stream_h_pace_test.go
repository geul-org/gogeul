//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHPace — going on saves nothing; a stop saves the pass position; a failing save is the error returned
package pack

import (
	"errors"
	"testing"
)

func TestStreamHPace(t *testing.T) {
	s, p := streamTestHPass(t, nil)
	if err := s.hPace(p); err != nil {
		t.Fatal(err)
	}
	stop := errors.New("stop")
	s.o.Pace = func(string) error { return stop }
	p.leaf = append(p.leaf, streamKV{Key: "k"})
	if err := s.hPace(p); !errors.Is(err, stop) || len(s.st.Close.HLeaf) != 1 {
		t.Fatalf("%v, saved leaf %v", err, s.st.Close.HLeaf)
	}
	p.heads.Close()
	p.w.WriteByte(1)
	if err := s.hPace(p); err == nil || errors.Is(err, stop) {
		t.Fatalf("failing save: %v", err)
	}
}
