//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirLatest — the newest live version through versions.geul; a missing list, a broken list and an all-deprecated list fail
package resolve

import (
	"bytes"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirLatest(t *testing.T) {
	f := newTypedFixture(t)
	if b, err := f.d.latest(sampleKeyOf(7)); err != nil || !bytes.Equal(b, f.v2) {
		t.Fatal(len(b), err)
	}
	if _, err := f.d.latest(sampleKeyOf(8)); !errors.Is(err, ErrNotFound) {
		t.Fatal("missing:", err)
	}
	f.d.Deprecate(f.g, []core.Word{0xffff, 0, 0, 7}, 1)
	f.d.Deprecate(f.g, []core.Word{0xffff, 0, 0, 7}, 2)
	if _, err := f.d.latest(sampleKeyOf(7)); !errors.Is(err, ErrNotFound) {
		t.Fatal("deprecated:", err)
	}
	f.d.write(listPath(sampleKeyOf(7)), []byte{1})
	if _, err := f.d.latest(sampleKeyOf(7)); err == nil {
		t.Fatal("broken: want error")
	}
}
