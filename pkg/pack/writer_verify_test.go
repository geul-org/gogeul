//ff:func feature=pack type=writer control=sequence
//ff:what TestWriterVerify — nothing to do without Verify or Check; an unopenable manifest, path checks and Check errors fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestWriterVerify(t *testing.T) {
	_, out := built(t, small())
	name := opened(t, out).ManifestName()
	w := testWriter(t, out, Batch{})
	w.o.Verify, w.o.Check = false, nil
	if err := w.verify("nonexistent"); err != nil {
		t.Fatal(err)
	}
	w.o.Verify = true
	if err := w.verify("nonexistent"); err == nil {
		t.Fatal("open")
	}
	w.sums = map[string][32]byte{"issuers/0x00000001.geul": {}}
	if err := w.verify(name); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	called := false
	w.o.Verify, w.o.Check = false, func(ctx context.Context, r *Reader) error { called = r != nil; return nil }
	if err := w.verify(name); err != nil || !called {
		t.Fatal(err)
	}
}
