//ff:func feature=cli type=helper control=sequence
//ff:what harness.metered — test fixture: run a command on a fresh pack archive over a meter; its error and the range reads it took
package cli

import (
	"bytes"
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

// metered opens the set in dir with empty caches, gives it to one command
// line, and counts from there.
func (h *harness) metered(dir string, args ...string) (int, error) {
	h.t.Helper()
	m := &meterStore{Store: pack.FileStore{Root: dir}}
	arch, err := pack.OpenRemote(context.Background(), m, pack.Options{})
	if err != nil {
		h.t.Fatal(err)
	}
	m.reads = 0
	a := h.app(&bytes.Buffer{})
	a.arch = arch
	cmd := newRootCmd(a)
	cmd.SetArgs(args)
	err = cmd.Execute()
	return m.reads, err
}
