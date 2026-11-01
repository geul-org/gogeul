//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamPace — no hook goes on; the hook gets the phase; its error comes back wrapped as ErrStopped
package pack

import (
	"errors"
	"testing"
)

func TestStreamPace(t *testing.T) {
	s := &Stream{}
	if err := s.pace("x"); err != nil {
		t.Fatal(err)
	}
	var got string
	s.o.Pace = func(p string) error { got = p; return nil }
	if err := s.pace("close:p"); err != nil || got != "close:p" {
		t.Fatalf("%q %v", got, err)
	}
	stop := errors.New("stop")
	s.o.Pace = func(string) error { return stop }
	if err := s.pace("close:p"); !errors.Is(err, ErrStopped) || !errors.Is(err, stop) {
		t.Fatal(err)
	}
}
