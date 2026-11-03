//ff:func feature=pack type=resolver control=sequence
//ff:what TestOpenRemote — one read-only set over any Store; an unreadable set fails
package pack

import (
	"context"
	"testing"
)

func TestOpenRemote(t *testing.T) {
	_, out := built(t, small())
	a, err := OpenRemote(context.Background(), FileStore{Root: out}, Options{})
	if err != nil || len(a.sets) != 1 || a.dir != "" {
		t.Fatal(err)
	}
	if b, err := a.ReadFile("issuers/0x00000001.geul"); err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	if _, err := OpenRemote(context.Background(), FileStore{Root: t.TempDir()}, Options{}); err == nil {
		t.Fatal("empty")
	}
}
