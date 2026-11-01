//ff:func feature=pack type=helper control=sequence
//ff:what openMetered — test fixture: the set in dir as a fresh Archive over a meterStore, the count starting after open
package pack

import (
	"context"
	"testing"
)

// openMetered opens dir with empty block caches.
func openMetered(t *testing.T, dir string) metered {
	t.Helper()
	m := &meterStore{Store: FileStore{Root: dir}}
	a, err := OpenRemote(context.Background(), m, Options{})
	if err != nil {
		t.Fatal(err)
	}
	m.reads = 0
	return metered{a: a, m: m}
}
