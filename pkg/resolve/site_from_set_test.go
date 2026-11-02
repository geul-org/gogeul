//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteFromSet — a set's path, ErrNotFound for no such set or path, the error of a broken set
package resolve

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSiteFromSet(t *testing.T) {
	var _ Source = (*pack.Reader)(nil)
	var _ Store = (*pack.Archive)(nil)
	ctx := context.Background()
	s := localSite(packedRoot(t, formsTree(t), false, nil))
	if b, err := s.fromSet(ctx, "0x0000", listPath(sampleKeyOf(7))); err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	if _, err := s.fromSet(ctx, "0x0000", listPath(sampleKeyOf(9))); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := s.fromSet(ctx, "0x0009", listPath(sampleKeyOf(7))); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	broken := &site{objects: pack.FileStore{Root: "/dev/null"}, sets: map[string]*pack.Reader{}}
	if _, err := broken.fromSet(ctx, "", "x"); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
