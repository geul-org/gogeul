//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestNewRootCmd — every planned command is there and the gone name command is not; a bad --format fails before any command runs
package cli

import (
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	root := newRootCmd(&app{})
	for _, name := range []string{"env", "ls", "show", "cat", "hash", "diff", "check", "blame", "grep", "resolve", "publish", "delta", "deprecate", "tombstone", "new", "open", "add", "def", "set", "rm", "meta", "issuer", "pack"} {
		if c, _, err := root.Find([]string{name}); err != nil || c.Name() != name {
			t.Errorf("%s missing", name)
		}
	}
	if c, _, err := root.Find([]string{"name"}); err == nil && c.Name() == "name" {
		t.Error("name: the document name command is gone")
	}
	h := newHarness(t)
	h.must("env")
	if _, err := h.run("", "env", "--format", "xml"); err == nil {
		t.Fatal("bad format: want error")
	}
}
